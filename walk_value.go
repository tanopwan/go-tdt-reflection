package main

import (
	"fmt"
	"reflect"
)

type domainValue struct {
	d reflect.Value
}

var values = []domainValue{
	{
		reflect.ValueOf(&Shape{
			Color: "Yellow",
			Edge:  4,
		}),
	},
	{
		reflect.ValueOf(&MyShape{
			Foos:  []Foo{{FooString: "foo1", Another: 1, Barfs: []Bar{{Barf: "B1"}}}, {FooString: "foo2", Another: 2}},
			Shape: Shape{Color: "Pink", Edge: 11},
		}),
	},
}

var valuePaths []string

func listValue() {
	for _, tt := range values {
		indirect := tt.d
		walkValue("object", indirect.Elem())
	}

	fmt.Printf("Result\n")
	for _, p := range valuePaths {
		fmt.Printf("%s\n", p)
	}
}

func walkValue(path string, v reflect.Value) {
	fmt.Printf("[%s]value: %v, type: %s, numfield: %d\n", path, v, v.Type(), v.NumField())

	for i := 0; i < v.NumField(); i++ {
		vv := v.Field(i)
		sf := v.Type().Field(i)
		fmt.Printf("[%s][%d]field type: %v, kind: %s\n", path, i, vv.Type(), vv.Kind())
		fmt.Printf("[%s][%d]struct field: %v, Tag: %s\n", path, i, sf, sf.Tag)

		if vv.Kind() == reflect.Struct {
			fmt.Printf("[%s][%d]Found struct\n", path, i)
			walkValue(path+"."+sf.Name, vv)
			return
		} else if vv.Kind() == reflect.Slice {
			fmt.Printf("[%s][%d]Found slice of %d\n", path, i, vv.Len())

			for ii := 0; ii < vv.Len(); ii++ {
				walkValue(fmt.Sprintf("%s.%s.%d", path, sf.Name, ii), vv.Index(ii))
			}
			return
		} else if vv.Kind() == reflect.String {
			fmt.Printf("[%s][%d][Test] String\n", path, i)
			valuePaths = append(valuePaths, path+"."+sf.Name)
		} else if vv.Kind() == reflect.Int {
			fmt.Printf("[%s][%d][Test] Int\n", path, i)
			valuePaths = append(valuePaths, path+"."+sf.Name)
		} else {
			fmt.Printf("[%s][%d][Test] Unknown\n", path, i)
			return
		}
	}
}
