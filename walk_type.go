package main

import (
	"fmt"
	"reflect"
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
}

// NewWalker ...
func NewWalker() Walker {
	return &walker{}
}

func (w *walker) Init(d []Domain) {
	w.domains = d
	w.listStruct()
}

func (w *walker) listStruct() {
	for _, domain := range w.domains {
		typeOf := reflect.TypeOf(domain.d)
		w.walkStruct(typeOf, typeOf.Name())
	}
}

func (w *walker) walkStruct(tt reflect.Type, path string) {
	for i := 0; i < tt.NumField(); i++ {
		f := tt.Field(i)
		switch f.Type.Kind() {
		case reflect.Struct:
			fmt.Printf("[struct] path: %s\n", fmt.Sprintf("%s.%s", path, f.Type.Name()))
			w.walkStruct(f.Type, fmt.Sprintf("%s.%s", path, f.Type.Name()))
		case reflect.Slice:
			fmt.Printf("[slice ] path: %s\n", fmt.Sprintf("%s.[]%s", path, f.Type.Elem().Name()))
			w.walkStruct(f.Type.Elem(), fmt.Sprintf("%s.[]%s", path, f.Type.Elem().Name()))
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
