package main

import (
	"fmt"
	"reflect"
)

type domain struct {
	d interface{}
}

var domains = []domain{
	{
		Shape{
			Color: "Yellow",
			Edge:  4,
		},
	},
	{
		MyShape{
			Foos:  []Foo{{FooString: "foo1", Another: 1, Barfs: []Bar{{Barf: "B1"}}}, {FooString: "foo2", Another: 2}},
			Shape: Shape{Color: "Pink", Edge: 11},
		},
	},
}

var paths []string

func listStruct() {
	for _, domain := range domains {
		typeOf := reflect.TypeOf(domain.d)
		dive(typeOf, typeOf.Name())
	}
}

func dive(tt reflect.Type, path string) {
	for i := 0; i < tt.NumField(); i++ {
		f := tt.Field(i)
		switch f.Type.Kind() {
		case reflect.Struct:
			fmt.Printf("[struct] path: %s\n", fmt.Sprintf("%s.%s", path, f.Type.Name()))
			dive(f.Type, fmt.Sprintf("%s.%s", path, f.Type.Name()))
		case reflect.Slice:
			fmt.Printf("[slice ] path: %s\n", fmt.Sprintf("%s.[]%s", path, f.Type.Elem().Name()))
			dive(f.Type.Elem(), fmt.Sprintf("%s.[]%s", path, f.Type.Elem().Name()))
		case reflect.String:
			res := fmt.Sprintf("%s.%s", path, f.Name)
			fmt.Printf("[string] path: %s\n", res)

			paths = append(paths, res)
			continue
		case reflect.Int:
			res := fmt.Sprintf("%s.%s", path, f.Name)
			fmt.Printf("[int   ] path: %s\n", res)
			paths = append(paths, res)
			continue
		default:
			fmt.Printf("[unknown] path: %s\n", fmt.Sprintf("%s.%s", path, f.Name))
		}
	}
}
