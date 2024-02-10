package main

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/dunhamsteve/iwork/index"
	"github.com/dunhamsteve/iwork/proto/KN"
	"github.com/dunhamsteve/iwork/proto/TN"
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
	imgs   map[string]uint64
	ix     *index.Index
	zr     *zip.ReadCloser
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
	ctx.imgs["Data/"+src] = dataId
	// not sure if this is px or pt.  It's px on the html side.
	width := fmt.Sprintf("%f", *image.OriginalSize.Width)
	height := fmt.Sprintf("%f", *image.OriginalSize.Height)
	return E("img", []string{"src", "", "width", width, "height", height, "class", "img_" + fmt.Sprint(dataId)})
}

var LE = binary.LittleEndian

func popcount(v uint16) int {
	var c int
	for ; v != 0; c++ {
		v &= v - 1
	}
	return c
}

func (ctx *Context) processTable(tm *TST.TableModelArchive) *html.Node {
	stringTable := ctx.ix.Deref(tm.DataStore.StringTable).(*TST.TableDataList).Entries
	richTable := ctx.ix.Deref(tm.DataStore.RichTextPayloadTable).(*TST.TableDataList).Entries

	// rc := *tm.NumberOfRows
	cc := *tm.NumberOfColumns

	// I found some hints at http://stingrayreader.sourceforge.net/workbook/numbers_13.html about how
	// this works, but I'm still flying blind.

	// so for now we assume at most one tile per row, and rows are in the right order.  I suspect long rows (more than
	// 255 columns) will have multiple tiles, however.  This would likely only happen in a spreadsheet.

	table := E("table")
	for _, tinfo := range tm.DataStore.Tiles.Tiles {
		tile := ctx.ix.Deref(tinfo.Tile).(*TST.Tile)
		for r, rinfo := range tile.RowInfos {
			tr := E("tr")
			table.AppendChild(tr)

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

				var cellType int
				// this has changed since I first wrote the code.  There is now a 4 in the first byte and the type in the next
				// the "stingrayreader" site says there is a halfword "version" and then the type, which I think worked at one
				// point, but I no longer have the file.
				if rinfo.CellStorageBuffer[offset] == 4 {
					cellType = int(rinfo.CellStorageBuffer[offset+1])
				} else {
					panic("not four")
					cellType = int(rinfo.CellStorageBuffer[offset+2])
				}

				// As far as I can tell, the records are variable length, with the pointer into the string/rich table at
				// the end, but this field seems to contain one bit per uint32 before the pointer to the string table
				// I suspect they are flags indicating which numbers/fields follow.
				flags := LE.Uint16(rinfo.CellStorageBuffer[offset+4 : offset+6])
				o := popcount(flags)*4 + 8 + int(offset)

				key := LE.Uint32(rinfo.CellStorageBuffer[o : o+4])

				// fmt.Printf("P %d %x %d %x\n", cellType, flags, popcount(flags), rinfo.CellStorageBuffer[o:o+4])
				// version := LE.Uint16(rinfo.CellStorageBuffer[offset : offset+2])
				// fmt.Println("XXX", c, version, cellType, hex.EncodeToString(rinfo.CellStorageBuffer[offset:]))
				switch cellType {
				case 0:
					// blank cells are type 0
				case 2: // number
					value := math.Float64frombits(LE.Uint64(rinfo.CellStorageBuffer[o : o+8]))
					td.AppendChild(E("p", fmt.Sprint(value)))
				case 5: // date
					value := math.Float64frombits(LE.Uint64(rinfo.CellStorageBuffer[o : o+8]))
					value += 978307200 // Apple to unix epoch
					tm := time.Unix(int64(value), 0)
					// We'll probably want to figure out formatting here.
					td.AppendChild(E("p", fmt.Sprint(tm)))
				case 6: // boolean
					value := math.Float64frombits(LE.Uint64(rinfo.CellStorageBuffer[o : o+8]))
					var label = "???"
					if value == 0 {
						label = "FALSE"
					} else if value == 0xf03f {
						label = "TRUE"
					}
					td.AppendChild(E("p", label))
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
				default:
					fmt.Printf("P %d %x %d %x\n", cellType, flags, popcount(flags), rinfo.CellStorageBuffer[o:o+8])
					fmt.Printf("CELL %d:%d type %d %s\n", r, c, cellType, hex.EncodeToString(rinfo.CellStorageBuffer[offset:]))
					td.AppendChild(E("p", fmt.Sprintf("UNKNOWN CELL TYPE %d", cellType)))
				}
			}
		}
	}
	rval := E("div")
	if tm.TableName != nil {
		// rval.AppendChild(E("h3", *tm.TableName)) // no need to show table name
	}
	rval.AppendChild(table)
	return rval
}

func (ctx *Context) processDrawable(ref *TSP.Reference) *html.Node {
	item := ctx.ix.Deref(ref)
	switch item.(type) {
	case *TSD.ImageArchive:
		return ctx.processImage(item.(*TSD.ImageArchive))
	case *TST.WPTableInfoArchive:
		table := item.(*TST.WPTableInfoArchive)
		tm := ctx.ix.Deref(table.Super.TableModel).(*TST.TableModelArchive)
		return ctx.processTable(tm)
	case *TST.TableInfoArchive:
		tm := ctx.ix.Deref(item.(*TST.TableInfoArchive).TableModel).(*TST.TableModelArchive)
		return ctx.processTable(tm)
	case *TSWP.ShapeInfoArchive:
		return ctx.processShapeInfo(item.(*TSWP.ShapeInfoArchive))
	case *TSD.GroupArchive:
		return ctx.processDrawableArchive(item.(*TSD.GroupArchive).Super)
	case *KN.PlaceholderArchive:
		return ctx.processShapeInfo(item.(*KN.PlaceholderArchive).Super)
	default:
		msg := fmt.Sprintf("*** Unhandled attachment type %T\n", item)
		fmt.Println(msg)
		return E("div", msg)
	}
	return nil
}

func (ctx *Context) processShapeInfo(sia *TSWP.ShapeInfoArchive) *html.Node {
	if cs := ctx.ix.Deref(sia.ContainedStorage).(*TSWP.StorageArchive); cs != nil {
		div := E("div")
		if ctx.storageToNode(cs, div) == nil {
			return div
		}
	}
	return ctx.processDrawableArchive(sia.Super.Super)
}

func (ctx *Context) processDrawableArchive(da *TSD.DrawableArchive) *html.Node {
	// do nothing...
	return nil
}

// storageToNode populates a html node with the contents of a StorageArchive. This happens with both the
// main body of the document and rich text table cells.
func (ctx *Context) storageToNode(bs *TSWP.StorageArchive, body *html.Node) error {
	ix := ctx.ix
	texts := bs.Text

	if len(texts) != 1 {
		return fmt.Errorf("FIXME - Expecting exactly one text, got %d", len(texts))
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
			switch ctx.ix.Deref(entry.Object).(type) {
			case *TSWP.DrawableAttachmentArchive:
				archive := ctx.ix.Deref(entry.Object).(*TSWP.DrawableAttachmentArchive)
				node := ctx.processDrawable(archive.Drawable)
				if node != nil {
					attachments = append(attachments, Attachment{pos, node})
				}
			case *TSWP.NumberAttachmentArchive:
				// do nothing...
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
	ctx.imgs = make(map[string]uint64)
	ctx.ix, err = index.Open(os.Args[1])
	must(err)
	ctx.zr, err = zip.OpenReader(os.Args[1])
	must(err)
	defer ctx.zr.Close()

	fmt.Println("Read", len(ctx.ix.Records), "records")

	var doc *html.Node
	switch ctx.ix.Type {
	case "pages":
		doc = ctx.processPages()
	case "numbers":
		doc = ctx.processNumbers()
	case "key":
		doc = ctx.processKeynote()
	}

	if len(os.Args) > 2 {
		w, err := os.Create(os.Args[2])
		must(err)
		defer w.Close()
		fmt.Println("Writing", os.Args[2])
		if strings.HasSuffix(os.Args[2], ".json") {
			out, err := json.MarshalIndent(ctx.ix, "", "  ")
			must(err)
			_, err = w.Write(out)
			must(err)

		} else {
			html.Render(w, doc)
		}

	} else {
		// html.Render(os.Stdout, doc)
	}
}

func (ctx *Context) renderImgData() *html.Node {
	if len(ctx.imgs) <= 0 {
		return nil
	}

	s := ""
	for _, f := range ctx.zr.File {
		if id, ok := ctx.imgs[f.Name]; ok {
			rc, _ := f.Open()
			defer rc.Close()

			imageBytes, _ := io.ReadAll(rc)
			s += "document.querySelectorAll('.img_" + fmt.Sprint(id) + "').forEach(function(e) {e.src='data:image/" + filepath.Ext(f.Name)[1:] + ";base64," + base64.StdEncoding.EncodeToString(imageBytes) + "';});"
		}
	}

	script := E("script")
	script.AppendChild(T(s))
	return script
}

// processPages translates a pages file.
func (ctx *Context) processPages() *html.Node {
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

	if img := ctx.renderImgData(); img != nil {
		body.AppendChild(img)
	}

	style := E("style")
	style.AppendChild(T("\np { margin: 0; }\n")) // reset paragraphs
	for k, v := range ctx.styles {
		style.AppendChild(T(fmt.Sprintf(".%s {\n%s}\n", k, v)))
	}
	head.AppendChild(style)
	return doc
}

func (ctx *Context) processNumbers() *html.Node {
	// Root of output document
	head, body := E("head", "\n", E("meta", []string{"charset", "utf-8"}), "\n"), E("body", "\n")
	doc := E("", E("html"), "\n", E("html", head, "\n", body))
	doc.Type = html.DocumentNode
	doc.FirstChild.Type = html.DoctypeNode
	da := ctx.ix.Records[1].(*TN.DocumentArchive)

	for _, ref := range da.Sheets {
		sheet := ctx.ix.Deref(ref).(*TN.SheetArchive)
		section := E("section", E("h2", "Sheet - ", *sheet.Name))
		doc.AppendChild(section)
		for _, ref := range sheet.DrawableInfos {
			// if this cast throws there are other kinds of drawables...
			section.AppendChild(ctx.processDrawable(ref))
		}
	}

	if img := ctx.renderImgData(); img != nil {
		body.AppendChild(img)
	}

	style := E("style")
	style.AppendChild(T("\np { margin: 0; }\n")) // reset paragraphs
	for k, v := range ctx.styles {
		style.AppendChild(T(fmt.Sprintf(".%s {\n%s}\n", k, v)))
	}
	head.AppendChild(style)

	return doc
}

// processKeynote translates a keynote file.
func (ctx *Context) processKeynote() *html.Node {
	// Root of output document
	head, body := E("head", "\n", E("meta", []string{"charset", "utf-8"}), "\n"), E("body", "\n")
	doc := E("", E("html"), "\n", E("html", head, "\n", body))
	doc.Type = html.DocumentNode
	doc.FirstChild.Type = html.DoctypeNode

	meta := ctx.ix.Records[2].(*TSP.PackageMetadata)
	ids := []uint64{}
	for _, comp := range meta.Components {
		if *comp.PreferredLocator == "Slide" {
			ids = append(ids, *comp.Identifier)
		}
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	container := E("container", []string{"class", "slide-container"})
	for _, id := range ids {
		slide := ctx.ix.Records[id].(*KN.SlideArchive)
		div := E("div", []string{"class", "slide"})
		for _, d := range append([]*TSP.Reference{slide.BodyPlaceholder}, slide.Drawables...) {
			if d == nil {
				continue
			}
			e := ctx.processDrawable(d)
			if e != nil {
				div.AppendChild(e)
			}
		}
		container.AppendChild(div)
	}
	body.AppendChild(container)

	if img := ctx.renderImgData(); img != nil {
		body.AppendChild(img)
	}

	style := E("style")
	style.AppendChild(T(".slide-container { display: flex; flex-direction: column; height: 100%;}\n.slide { flex: 1; background-color: #f0f0f0; padding: 50px; font-size: 24px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.3); transition: transform 0.5s ease, box-shadow 0.5s ease;}\n.slide:hover { transform: scale(1.01); box-shadow: 0 0 20px rgba(0, 0, 0, 0.5);}\np { margin: 0; }\n")) // reset paragraphs
	for k, v := range ctx.styles {
		style.AppendChild(T(fmt.Sprintf(".%s {\n%s}\n", k, v)))
	}
	head.AppendChild(style)
	return doc
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

// Helper functions for debugging

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
