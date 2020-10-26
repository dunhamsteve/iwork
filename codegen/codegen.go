package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

var foo = `
package index
import (
    "errors"
    "github.com/golang/protobuf/proto"
    "fmt"
)

func decode(typ uint32, payload []byte) (interface{}, error) {
    switch typ {
        {{range $key, $value := .}}
        case {{$key}}:
            var value = &{{$value}}{}
            err := proto.Unmarshal(payload, value)
            return value, err
        {{end}}
        default:
            return nil,errors.New(fmt.Sprintf("Unknown type %d", typ))
    }
}
`

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	must(err)
	var info map[string]string

	must(json.Unmarshal(data, &info))

	tmpl, err := template.New("test").Parse(foo)
	must(err)

	must(tmpl.Execute(os.Stdout, info))
}
