package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Domain ...
type Domain struct {
	d interface{}
}

// Walker ...
type Walker interface {
	Init(d []Domain)
}

type walker struct {
	domains []Domain
	paths   []string
	results []ResultRow
}

// ResultRow ...
type ResultRow struct {
	path string
	body string
	tag  string
}

// NewWalker ...
func NewWalker() Walker {
	return &walker{}
}

func (w *walker) Init(d []Domain) {
	w.domains = d
	w.listStruct()
	w.generateTable()
}

func (w *walker) listStruct() {
	for i, domain := range w.domains {
		typeOf := reflect.TypeOf(domain.d)
		w.walkStruct(typeOf, fmt.Sprintf("%d.%s", i, typeOf.Name()))
	}
}

func (w *walker) walkStruct(tt reflect.Type, path string) {
	for i := 0; i < tt.NumField(); i++ {
		f := tt.Field(i)
		switch f.Type.Kind() {
		case reflect.Struct:
			fmt.Printf("[struct] path: %s\n", fmt.Sprintf("%s.%s", path, f.Name))
			w.walkStruct(f.Type, fmt.Sprintf("%s.%s", path, f.Name))
		case reflect.Slice:
			fmt.Printf("[slice ] path: %s\n", fmt.Sprintf("%s.[]%s", path, f.Name))
			w.walkStruct(f.Type.Elem(), fmt.Sprintf("%s.[]%s", path, f.Name))
		case reflect.String:
			res := fmt.Sprintf("%s.%s", path, f.Name)
			fmt.Printf("[string] path: %s\n", res)

			w.paths = append(w.paths, res)
			continue
		case reflect.Int:
			res := fmt.Sprintf("%s.%s", path, f.Name)
			fmt.Printf("[int   ] path: %s\n", res)
			w.paths = append(w.paths, res)
			continue
		default:
			fmt.Printf("[unknown] path: %s\n", fmt.Sprintf("%s.%s", path, f.Name))
		}
	}
}

func (w *walker) generateTable() {
	for _, path := range w.paths {
		nodes := strings.Split(path, ".")
		c, _ := strconv.Atoi(nodes[0])
		p := reflect.ValueOf(&w.domains[c].d)
		value := p.Elem()
		fmt.Printf("[path] %s - value of domain: %v, with type: %s, typeOf: %s\n", path, value, value.Type(), reflect.TypeOf(value.Interface()))

		t := reflect.TypeOf(value.Interface())
		fmt.Printf("%v", t.Field(0))

		v := value
		if len(nodes) > 2 {
			for i := 2; i < len(nodes); i++ {
				fmt.Printf(">> level: %d, node: %s\n", i, nodes[i])
				var vv reflect.Value
				if strings.HasPrefix(nodes[i], "[]") {
					fmt.Printf(">> slice\n")
					// Slice
					name := strings.TrimPrefix(nodes[i], "[]")
					if t.FieldByName(name).Len() == 0 {
						continue
					}

					vv = t.FieldByName(name).Index(0)
				} else {
					vv = t.FieldByName(nodes[i])
				}

				fmt.Printf(">> value of field: %v, with type: %s\n", vv, vv.Type())
				v = vv
			}
		}

		if v.Kind() == reflect.String {
			v.SetString("")
		}

		// payload, err := json.Marshal(newRoot.Interface())
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Printf("[%s][Test] request payload: %s\n", path, string(payload))
	}
}

func testStringTags(tags []string) {

}
