package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dunhamsteve/iwork/index"
	"github.com/dunhamsteve/iwork/proto/TP"
	"github.com/dunhamsteve/iwork/proto/TSD"
	"github.com/dunhamsteve/iwork/proto/TSP"
	"github.com/dunhamsteve/iwork/proto/TST"
	"github.com/dunhamsteve/iwork/proto/TSWP"

	"golang.org/x/net/html"
)

// T is a helper function for building html text nodes.
func T(value string) *html.Node {
	return &html.Node{Type: html.TextNode, Data: value}
}

// E is a helper function for building HTML nodes.
func E(tag string, children ...interface{}) *html.Node {
	node := &html.Node{}
	node.Data = tag
	node.Type = html.ElementNode

	for _, child := range children {
		switch child.(type) {
		case string:
			node.AppendChild(T(child.(string)))
		case *html.Node:
			node.AppendChild(child.(*html.Node))
		case []string:
			args := child.([]string)
			for i := 0; i < len(args)-1; i += 2 {
				node.Attr = append(node.Attr, html.Attribute{Key: args[i], Val: args[i+1]})
			}
		default:
			log.Fatalf("Unhandled type %T in E", child)
		}
	}
	return node
}

type Context struct {
	styles map[string]string
	ix     *index.Index
}

type Attachment struct {
	pos  uint32
	node *html.Node
}

func (ctx *Context) processImage(image *TSD.ImageArchive) *html.Node {
	dataId := *image.Data.Identifier
	meta := ctx.ix.Records[2].(*TSP.PackageMetadata)
	var src string
	for _, data := range meta.Datas {
		if dataId == *data.Identifier {
			if data.FileName != nil {
				src = *data.FileName
			} else {
				fmt.Println("No filename: %#v\n", data)
				src = *data.PreferredFileName
			}
		}
	}
	// not sure if this is px or pt.  It's px on the html side.
	width := fmt.Sprintf("%f", *image.OriginalSize.Width)
	height := fmt.Sprintf("%f", *image.OriginalSize.Height)
	return E("img", []string{"src", src, "width", width, "height", height})
}

var LE = binary.LittleEndian

func popcount(v uint16) int {
	var c int
	for ; v != 0; c++ {
		v &= v - 1
	}
	return c
}

func (ctx *Context) processTable(table *TST.WPTableInfoArchive) *html.Node {
	tm := ctx.ix.Deref(table.Super.TableModel).(*TST.TableModelArchive)

	stringTable := ctx.ix.Deref(tm.DataStore.StringTable).(*TST.TableDataList).Entries
	richTable := ctx.ix.Deref(tm.DataStore.RichTextPayloadTable).(*TST.TableDataList).Entries

	// rc := *tm.NumberOfRows
	cc := *tm.NumberOfColumns

	// I found some hints at http://stingrayreader.sourceforge.net/workbook/numbers_13.html about how
	// this works, but I'm still flying blind.

	// so for now we assume at most one tile per row, and rows are in the right order.  I suspect long rows (more than
	// 255 columns) will have multiple tiles, however.  This would likely only happen in a spreadsheet.

	rval := E("table")
	for _, tinfo := range tm.DataStore.Tiles.Tiles {
		tile := ctx.ix.Deref(tinfo.Tile).(*TST.Tile)
		for _, rinfo := range tile.RowInfos {
			tr := E("tr")
			rval.AppendChild(tr)

			offsets := make([]uint16, len(rinfo.CellOffsets)/2)
			binary.Read(bytes.NewBuffer(rinfo.CellOffsets), LE, offsets)
			// Naïvely assuming that the index is the column number, per the stringrayreader code.
			// FIXME - figure out the right way to determine column number.
			for c, offset := range offsets {
				if uint32(c) >= cc {
					break
				}
				td := E("td")
				tr.AppendChild(td)

				// 0xffff is an empty cell (This only occurs at the end in my sample document.)
				if offset == 65535 {
					continue
				}

				// version := LE.Uint16(rinfo.CellStorageBuffer[offset : offset+2])
				cellType := int(rinfo.CellStorageBuffer[offset+2])

				// As far as I can tell, the records are variable length, with the pointer into the string/rich table at
				// the end, but this field seems to contain one bit per uint32 before the pointer to the string table
				// I suspect they are flags indicating which numbers/fields follow.
				flags := LE.Uint16(rinfo.CellStorageBuffer[offset+4 : offset+6])
				o := popcount(flags)*4 + 8 + int(offset)

				key := LE.Uint32(rinfo.CellStorageBuffer[o : o+4])

				// fmt.Printf("P %x %d %d %x\n", flags, popcount(flags), rinfo.CellStorageBuffer[o:o+4])
				// fmt.Println("XXX", c, version, cellType, hex.EncodeToString(rinfo.CellStorageBuffer[offset:]))
				switch cellType {
				case 3:
					for _, entry := range stringTable {
						if *entry.Key == key {
							td.AppendChild(E("p", *entry.String_))
						}
					}

				case 9:
					for _, entry := range richTable {
						if *entry.Key == key {
							rt := ctx.ix.Deref(entry.RichTextPayload).(*TST.RichTextPayloadArchive)
							st := ctx.ix.Deref(rt.Storage).(*TSWP.StorageArchive)
							ctx.storageToNode(st, td)
						}
					}
				}
			}
		}
	}
	return rval
}

func (ctx *Context) processAttachment(attach *TSWP.DrawableAttachmentArchive) *html.Node {
	item := ctx.ix.Deref(attach.Drawable)
	switch item.(type) {
	case *TSD.ImageArchive:
		return ctx.processImage(item.(*TSD.ImageArchive))
	case *TST.WPTableInfoArchive:
		return ctx.processTable(item.(*TST.WPTableInfoArchive))
	default:
		fmt.Printf("Unhandled attachment type %T\n", item)
	}
	return nil
}

// storageToNode populates a html node with the contents of a StorageArchive. This happens with both the
// main body of the document and rich text table cells.
func (ctx *Context) storageToNode(bs *TSWP.StorageArchive, body *html.Node) error {
	ix := ctx.ix
	texts := bs.Text

	if len(texts) != 1 {
		log.Printf("FIXME - Expecting exactly one text, got %d", len(texts))
	}
	text := texts[0]

	// Offsets are in terms of unicode runes, so we have to convert to runes
	rr := []rune(text)

	// <p>
	parStyles := bs.TableParaStyle.Entries

	var attachments []Attachment
	if bs.TableAttachment != nil {
		for _, entry := range bs.TableAttachment.Entries {
			pos := *entry.CharacterIndex
			archive := ctx.ix.Deref(entry.Object).(*TSWP.DrawableAttachmentArchive)
			node := ctx.processAttachment(archive)
			if node != nil {
				attachments = append(attachments, Attachment{pos, node})
			}
		}
	}
	// bs.TableListStyle - seems to change on headings, look into it.

	// A null style seems to imply "use the previous class," so this is declared outside the loop.
	var className string

	// build paragraphs
	for i, e := range parStyles {

		pos := *e.CharacterIndex
		end := uint32(len(rr))
		if i+1 < len(parStyles) {
			end = *parStyles[i+1].CharacterIndex
		}

		for len(attachments) > 0 && attachments[0].pos < end {
			if attachments[0].pos != pos {
				fmt.Printf("FIXME - attachment not at start of paragraph - pstart=%d pend=%d att=%d par=%#v\n",
					pos, end, attachments[0].pos, string(rr[pos:end]))
			}
			body.AppendChild(attachments[0].node)
			attachments = attachments[1:]
		}

		tag := "p"

		// Get style, change tag if appropriate.
		if e.Object != nil {
			ref := ix.Deref(e.Object).(*TSWP.ParagraphStyleArchive)
			className = fmt.Sprintf("ps%d", *e.Object.Identifier)

			// Some properties are inherited (e.g. if you apply a style and then tweak it.)
			// We can't just include both because FirstLineIndent in parent can combine with LeftIndent in child
			// to produce a css text-indent.
			if ref.Super.Parent != nil {
				parent := ix.Deref(ref.Super.Parent).(*TSWP.ParagraphStyleArchive)
				mergeCharProps(ref.CharProperties, parent.CharProperties)
				mergeParaProps(ref.ParaProperties, parent.ParaProperties)

				if parent.Super.Parent != nil {
					panic("Need recursion here")
				}
			}

			ctx.styles[className] = translateParaProps(ref.ParaProperties) + translateCharProps(ref.CharProperties)

			if ref.ParaProperties.OutlineLevel != nil {
				level := *ref.ParaProperties.OutlineLevel
				if level < 7 {
					tag = fmt.Sprintf("h%d", level)
				}
			}
		}

		p := E(tag, []string{"class", className})
		// <span> <em> and <b>
		if bs.TableCharStyle != nil {
			charStyles := bs.TableCharStyle.Entries
			for i, e := range charStyles { // build any span/em/b as needed
				cs := *e.CharacterIndex
				if cs < pos {
					continue
				}
				if cs >= end {
					break
				}
				ce := uint32(len(rr))
				if i+1 < len(charStyles) {
					ce = *charStyles[i+1].CharacterIndex
				}
				if ce > end {
					if e.Object != nil {
						fmt.Println("ERR? ce > end", ce, end, e.Object)
					}
					ce = end
				}
				if cs > pos {
					p.AppendChild(T(string(rr[pos:cs])))
					pos = cs
				}
				if e.Object != nil {
					ref := ix.Deref(e.Object).(*TSWP.CharacterStyleArchive)
					key := fmt.Sprintf("ss%d", *e.Object.Identifier)

					if ref.Super.Parent != nil {
						parent := ix.Deref(ref.Super.Parent).(*TSWP.CharacterStyleArchive)
						mergeCharProps(ref.CharProperties, parent.CharProperties)
						if parent.Super.Parent != nil {
							panic("Need recursion here")
						}
					}

					style := translateCharProps(ref.CharProperties)

					// convert em/bold to tags
					switch style {
					case "  font-style: italic;\n":
						p.AppendChild(E("em", string(rr[cs:ce])))
					case "  font-weight: bold;\n":
						p.AppendChild(E("b", string(rr[cs:ce])))
					default:
						ctx.styles[key] = style
						p.AppendChild(E("span", []string{"class", key}, string(rr[cs:ce])))
					}
				} else {
					p.AppendChild(T(string(rr[cs:ce])))
				}
				pos = ce
			}
		}
		p.AppendChild(T(string(rr[pos:end])))
		body.AppendChild(p)
		body.AppendChild(T("\n"))
	}

	return nil
}

type Style map[string]interface{}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf(`Converts pages files to html

Usage:
    %s infile.pages outfile.html

`, os.Args[0])
		return
	}

	fmt.Println("Processing", os.Args[1])

	var err error
	var ctx Context

	ctx.styles = make(map[string]string)
	ctx.ix, err = index.Open(os.Args[1])
	must(err)

	fmt.Println("Read", len(ctx.ix.Records), "records")

	// Root of output document
	head, body := E("head", "\n", E("meta", []string{"charset", "utf-8"}), "\n"), E("body", "\n")
	doc := E("", E("html"), "\n", E("html", head, "\n", body))
	doc.Type = html.DocumentNode
	doc.FirstChild.Type = html.DoctypeNode

	da := ctx.ix.Records[1].(*TP.DocumentArchive)
	bs := ctx.ix.Deref(da.BodyStorage).(*TSWP.StorageArchive)

	fda := ctx.ix.Deref(da.FloatingDrawables).(*TP.FloatingDrawablesArchive)
	if len(fda.PageGroups) != 0 {
		fmt.Println(`WARNING - 
            This document has floating drawables (e.g. floating images/tables/text blocks) which we don't handle in HTML
            conversion.
            
            Figuring out where to place them in the document would probably be tricky.
            
`)
	}

	ctx.storageToNode(bs, body)

	style := E("style")
	style.AppendChild(T("\np { margin: 0; }\n")) // reset paragraphs
	for k, v := range ctx.styles {
		style.AppendChild(T(fmt.Sprintf(".%s {\n%s}\n", k, v)))
	}
	head.AppendChild(style)

	if len(os.Args) > 2 {
		w, err := os.Create(os.Args[2])
		must(err)
		defer w.Close()
		fmt.Println("Writing", os.Args[2])
		html.Render(w, doc)
	} else {
		// html.Render(os.Stdout, doc)
	}
}

func mergeCharProps(props *TSWP.CharacterStylePropertiesArchive, parent *TSWP.CharacterStylePropertiesArchive) {
	if parent != nil {
		if props.Bold == nil {
			props.Bold = parent.Bold
		}
		if props.Italic == nil {
			props.Italic = parent.Italic
		}
		if props.FontSize == nil {
			props.FontSize = parent.FontSize
		}
		if props.FontName == nil {
			props.FontName = parent.FontName
		}
	}
}

// translateCharProps converts a TSWP.CharacterStylePropertiesArchive into CSS
func translateCharProps(props *TSWP.CharacterStylePropertiesArchive) string {

	rval := ""
	if props.Bold != nil && *props.Bold {
		rval += "  font-weight: bold;\n"
	}
	if props.Italic != nil && *props.Italic {
		rval += "  font-style: italic;\n"
	}
	if props.FontSize != nil {
		rval += fmt.Sprintf("  font-size: %fpt;\n", *props.FontSize)
	}
	if props.FontName != nil {
		rval += fmt.Sprintf("  font-family: '%s';\n", *props.FontName)
	}
	return rval
}

func mergeParaProps(props *TSWP.ParagraphStylePropertiesArchive, parent *TSWP.ParagraphStylePropertiesArchive) {
	if parent != nil {
		if props.LeftIndent == nil {
			props.LeftIndent = parent.LeftIndent
		}
		if props.RightIndent == nil {
			props.RightIndent = parent.RightIndent
		}
		if props.SpaceBefore == nil {
			props.SpaceBefore = parent.SpaceBefore
		}
		if props.SpaceAfter == nil {
			props.SpaceAfter = parent.SpaceAfter
		}
		if props.FirstLineIndent == nil {
			props.FirstLineIndent = parent.FirstLineIndent
		}
	}
}

// translateParaProps converts a TSWP.ParagraphStylePropertiesArchive into CSS.
func translateParaProps(props *TSWP.ParagraphStylePropertiesArchive) string {
	rval := ""

	if props.LeftIndent != nil && *props.LeftIndent != 0. {
		rval += fmt.Sprintf("  margin-left: %fpt;\n", *props.LeftIndent)
	}

	if props.FirstLineIndent != nil {
		textIndent := *props.FirstLineIndent
		if props.LeftIndent != nil {
			textIndent = textIndent - *props.LeftIndent
		}
		if textIndent != 0. {
			rval += fmt.Sprintf("  text-indent: %fpt;\n", textIndent)
		}
	}

	if props.RightIndent != nil && *props.RightIndent > 0. {
		rval += fmt.Sprintf("  margin-right: %fpt;\n", *props.RightIndent)
	}
	if props.SpaceBefore != nil && *props.SpaceBefore > 0. {
		rval += fmt.Sprintf("  margin-top: %fpt;\n", *props.SpaceBefore)
	}
	if props.SpaceAfter != nil && *props.SpaceAfter > 0. {
		rval += fmt.Sprintf("  margin-bottom: %fpt;\n", *props.SpaceAfter)
	}

	return rval
}

// Debugging crap

func must(err error) {
	if err != nil {
		panic(err)
	}
}
func writejson(foo interface{}, fn string) {
	a, err := json.MarshalIndent(foo, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	ioutil.WriteFile(fn, a, 0644)
}

func dumpjson(foo interface{}) {
	a, err := json.MarshalIndent(foo, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(a))
}

func dump(foo interface{}) {
	fmt.Printf("%#v\n", foo)
}

func ptype(x interface{}) {
	fmt.Printf("type %T\n", x)
}
