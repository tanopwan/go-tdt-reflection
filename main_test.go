package main_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	m "github.com/tanopwan/go-tdt-reflection"
)

type testDomain struct {
	requestDomain reflect.Value
	header        string
	path          string
}

var testDomains = []testDomain{
	{reflect.ValueOf(&m.Shape{
		Color: "Yellow",
		Edge:  4,
	}), "header", "/path/to/resource1"},
	{reflect.ValueOf(&m.MyShape{
		Shapes: []m.Shape{m.Shape{Color: "Blue", Edge: 3}, m.Shape{Color: "Purple", Edge: 2}},
		Shape:  m.Shape{},
	}), "header", "/path/to/resource2"},
}

var (
	tag       = "test"
	tag_empty = "empty"
	tag_dive  = "dive"
)

func TestDomains(t *testing.T) {
	for _, tt := range testDomains {
		indirect := tt.requestDomain
		t.Run(indirect.Elem().Type().Name(), func(t *testing.T) {
			dive(t, "", indirect.Elem())
		})
	}
}

func dive(t *testing.T, path string, indirectElem reflect.Value) {
	for i := 0; i < indirectElem.NumField(); i++ {
		newElem := reflect.New(indirectElem.Type())
		newElem.Elem().Set(reflect.ValueOf(indirectElem.Interface()))
		fmt.Printf(">>> newElem: %v, type: %s, no: %d\n", newElem, newElem.Type(), newElem.Elem().NumField())
		elem := newElem.Elem()

		t.Run(elem.Type().Field(i).Name, func(t *testing.T) {
			fmt.Printf("%sfield type: %v, kind: %s\n", path, elem.Field(i).Type(), elem.Field(i).Kind())
			fmt.Printf("%sstruct field: %v, Tag: %s\n", path, elem.Type().Field(i), elem.Type().Field(i).Tag)
			if elem.Field(i).Kind() == reflect.Struct {
				dive(t, path+"...", elem.Field(i))
			} else if elem.Field(i).Kind() == reflect.Slice {
				for ii := 0; ii < elem.Field(i).Len(); ii++ {

					indirectFieldIndex := elem.Field(i).Index(ii)
					dive(t, fmt.Sprintf("%s...[%d]", path, ii), indirectFieldIndex)
				}
			} else if elem.Field(i).Kind() == reflect.String {
				fmt.Printf("%s[Test] String\n", path)
				callTest(t, path, elem.Field(i).Addr())

				payload, err := json.Marshal(elem.Interface())
				if err != nil {
					panic(err)
				}
				fmt.Printf("%s[Test]... request payload: %s\n", path, string(payload))
			} else if elem.Field(i).Kind() == reflect.Int {
				fmt.Printf("%s[Test] Int\n", path)
				callTest(t, path, elem.Field(i).Addr())

				payload, err := json.Marshal(elem.Interface())
				if err != nil {
					panic(err)
				}
				fmt.Printf("%s[Test]... request payload: %s\n", path, string(payload))
			} else {
				fmt.Printf("%s[Test] Unknown\n", path)
			}
		})
	}
}

func callTest(t *testing.T, path string, indirect reflect.Value) {
	if indirect.Elem().Kind() == reflect.String {
		indirect.Elem().SetString("")
	} else if indirect.Elem().Kind() == reflect.Slice {
		indirect.Elem().SetLen(0)
	} else if indirect.Elem().Kind() == reflect.Int {
		indirect.Elem().SetInt(0)
	} else {
		fmt.Printf("%s[Test] ... skip field unknown Kind: %v\n", path, indirect.Elem().Kind())
	}
}
