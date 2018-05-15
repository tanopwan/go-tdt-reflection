package main_test

import (
	"fmt"
	"reflect"
	"testing"

	m "github.com/tanopwan/go-tdt-reflection"
)

type testDomain struct {
	requestDomain interface{}
	header        string
	path          string
}

var testDomains = []testDomain{
	{m.Shape{
		Color: "Yellow",
		Edge:  4,
	}, "header", "/path/to/resource1"},
	{m.MyShape{
		Shapes: []m.Shape{m.Shape{}, m.Shape{}},
		Shape:  m.Shape{},
	}, "header", "/path/to/resource2"},
}

var (
	tag       = "test"
	tag_empty = "empty"
	tag_dive  = "dive"
)

func TestDomains(t *testing.T) {
	for _, tt := range testDomains {
		reflectType := reflect.TypeOf(tt.requestDomain)
		fmt.Printf("reflectType: %s\n", reflectType)
		indirect := reflect.Indirect(reflect.ValueOf(tt.requestDomain))
		t.Run(reflectType.Name(), func(t *testing.T) {
			dive(t, "", indirect)
		})
	}
}

func dive(t *testing.T, path string, indirect reflect.Value) {
	for i := 0; i < indirect.NumField(); i++ {
		t.Run(indirect.Type().Field(i).Name, func(t *testing.T) {
			fmt.Printf("%sfield type: %v, kind: %s\n", path, indirect.Field(i).Type(), indirect.Field(i).Kind())
			fmt.Printf("%sstruct field: %v, Tag: %s\n", path, indirect.Type().Field(i), indirect.Type().Field(i).Tag)
			if indirect.Field(i).Kind() == reflect.Struct {
				dive(t, path+"...", indirect.Field(i))
			} else if indirect.Field(i).Kind() == reflect.Slice {
				for ii := 0; ii < indirect.Field(i).Len(); ii++ {
					dive(t, fmt.Sprintf("%s...[%d]", path, ii), indirect.Field(i).Index(ii))
				}
			} else if indirect.Field(i).Kind() == reflect.String {
				fmt.Printf("%s[Test] String\n", path)
			} else if indirect.Field(i).Kind() == reflect.Int {
				fmt.Printf("%s[Test] Int\n", path)
			} else {
				fmt.Printf("%s[Test] Unknown\n", path)
			}
		})
	}
}
