package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dunhamsteve/iwork/index"
	"github.com/dunhamsteve/iwork/proto/TP"
	"github.com/dunhamsteve/iwork/proto/TSWP"

	"golang.org/x/net/html"
)

// Helper function for building html text nodes
func T(value string) *html.Node {
	return &html.Node{Type: html.TextNode, Data: value}
}

// Helper function for building HTML nodes
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

func main() {
	if len(os.Args) < 3 {
		fmt.Printf(`Converts pages files to html

Usage:
    %s infile.pages outfile.html

`, os.Args[0])
		return
	}

	fmt.Println("Processing", os.Args[1])
	ix, err := index.Open(os.Args[1])
	must(err)

	fmt.Println(len(ix.Records), "records")

	da := ix.Records[1].(*TP.DocumentArchive)

	// This has most of what we want, aside from the actual style definitions.
	bs := ix.Deref(da.BodyStorage).(*TSWP.StorageArchive)
	texts := bs.Text

	if len(texts) != 1 {
		log.Printf("WARNING - Expecting exactly one text, got %d", len(texts))
	}
	text := texts[0]

	// Offsets are in terms of unicode runes, so we have to convert to runes
	rr := []rune(text)

	// We collect CSS styles here, as we reference them, so we don't write duplicates
	styleMap := make(map[string]string)

	// <p>
	parStyles := bs.TableParaStyle.Entries

	// <span> <em> and <b>
	charStyles := bs.TableCharStyle.Entries

	// bs.TableListStyle - seems to change on headings, look into it.

	// Root of output document
	head, body := E("head", "\n"), E("body", "\n")
	doc := E("html", head, "\n", body)

	var className string

	// build paragraphs
	for i, e := range parStyles {
		pos := *e.CharacterIndex
		end := uint32(len(rr))
		if i+1 < len(parStyles) {
			end = *parStyles[i+1].CharacterIndex
		}

		// Get style
		if e.Object != nil {
			ref := ix.Deref(e.Object).(*TSWP.ParagraphStyleArchive)
			className = fmt.Sprintf("ps%d", *e.Object.Identifier)
			styleMap[className] = translateParaProps(ref.ParaProperties) + translateCharProps(ref.CharProperties)
		}

		p := E("p", []string{"class", className})
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
				style := translateCharProps(ref.CharProperties)

				// convert em/bold to tags
				switch style {
				case "  font-style: italic;\n":
					p.AppendChild(E("em", string(rr[cs:ce])))
				case "  font-weight: bold;\n":
					p.AppendChild(E("b", string(rr[cs:ce])))
				default:
					styleMap[key] = style
					p.AppendChild(E("span", []string{"class", key}, string(rr[cs:ce])))
				}
			} else {
				p.AppendChild(T(string(rr[cs:ce])))
			}
			pos = ce
		}

		p.AppendChild(T(string(rr[pos:end])))
		body.AppendChild(p)
		// body.AppendChild(E("p", []string{ "class", className}, string(rr[pos:end])))
		body.AppendChild(T("\n"))
	}

	style := E("style")
	style.AppendChild(T("\np { margin: 0; }\n")) // reset paragraphs
	for k, v := range styleMap {
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
		rval += fmt.Sprintf("  font-size: %fpx;\n", *props.FontSize)
	}
	if props.FontName != nil {
		rval += fmt.Sprintf("  font-family: '%s';\n", *props.FontName)
	}
	return rval
}

// translateParaProps converts a TSWP.ParagraphStylePropertiesArchive into CSS
func translateParaProps(props *TSWP.ParagraphStylePropertiesArchive) string {
	rval := ""
	if props.FirstLineIndent != nil && *props.FirstLineIndent > 0. {
		rval += fmt.Sprintf("  text-indent: %fpx;\n", *props.FirstLineIndent)
	}
	if props.LeftIndent != nil && *props.LeftIndent > 0. {
		rval += fmt.Sprintf("  margin-left: %fpx;\n", *props.LeftIndent)
	}
	if props.RightIndent != nil && *props.RightIndent > 0. {
		rval += fmt.Sprintf("  margin-right: %fpx;\n", *props.RightIndent)
	}
	if props.SpaceBefore != nil && *props.SpaceBefore > 0. {
		rval += fmt.Sprintf("  margin-top: %fpx;\n", *props.SpaceBefore)
	}
	if props.SpaceAfter != nil && *props.SpaceAfter > 0. {
		rval += fmt.Sprintf("  margin-bottom: %fpx;\n", *props.SpaceAfter)
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
