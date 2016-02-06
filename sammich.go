package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var smap = `package {{.Package}}

import (
	"sync"{{range .Imports}}
	"{{.}}"
{{end}})

// {{.Name}}Map is a synchronized map[{{.KeyPtr}}{{.KeyType}}]{{.ValPtr}}{{.ValType}}.
//
// It is not sized by default.
type {{.Name}}Map struct {
	sync.RWMutex
	dict map[{{.KeyPtr}}{{.KeyType}}]{{.ValPtr}}{{.ValType}}
}

// New creates a new synchronized map of {{.KeyPtr}}{{.KeyType}} to {{.ValPtr}}{{.ValType}}.
func New{{.Name}}Map() *{{.Name}}Map {
	return &{{.Name}}Map{
		dict: map[{{.KeyPtr}}{{.KeyType}}]{{.ValPtr}}{{.ValType}}{},
	}
}

// Get takes a {{.KeyPtr}}{{.KeyType}} and returns a {{.ValPtr}}{{.ValType}} if
// found, and bool indicating whether it was found.
func (o *{{.Name}}Map) Get(k {{.KeyPtr}}{{.KeyType}}) ({{.ValPtr}}{{.ValType}}, bool) {
	o.RLock()
	v, ok := o.dict[k]
	o.RUnlock()
	return v, ok
}

func (o *{{.Name}}Map) Put(k {{.KeyPtr}}{{.KeyType}}, v {{.ValPtr}}{{.ValType}}) {
	o.Lock()
	o.dict[k] = v
	o.Unlock()
}

// Delete removes a value from the map.
//
// If the key is nil or not found, Delete is a no-op.
func (o *{{.Name}}Map) Delete(k {{.KeyPtr}}{{.KeyType}}) {
	o.Lock()
	delete(o.dict,k)
	o.Unlock()
}
`

// main runs the program.
//
// https://xkcd.com/149/
func main() {
	tt := template.Must(template.New("smap").Parse(smap))
	k := os.Args[2]
	kd := ""
	vd := ""
	if strings.HasPrefix(k, "*") {
		k = k[1:]
		kd = "*"
	}
	v := os.Args[3]
	if strings.HasPrefix(v, "*") {
		v = v[1:]
		vd = "*"
	}

	imports := []string{}
	if len(os.Args) > 4 {
		for i := 4; i < len(os.Args); i++ {
			imports = append(imports, os.Args[i])
		}
	}

	dest := fmt.Sprintf("%s_%s_map.go", k, v)
	file, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Could not create %s: %s\n", dest, err)
		os.Exit(1)
	}

	vals := map[string]interface{}{
		"Name":    strings.Title(k) + strings.Title(v),
		"KeyType": k,
		"ValType": v,
		"KeyPtr":  kd,
		"ValPtr":  vd,
		"Package": os.Getenv("GOPACKAGE"),
		"Imports": imports,
	}
	tt.Execute(file, vals)

	file.Close()
}
