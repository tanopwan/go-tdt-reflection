package main_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
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
		Shape:  m.Shape{Color: "Pink", Edge: 11},
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
			dive(t, "object", indirect.Elem(), indirect.Elem())
		})
	}
}

func dive(t *testing.T, path string, indirectElem reflect.Value, root reflect.Value) {
	for i := 0; i < indirectElem.NumField(); i++ {
		newElem := reflect.New(indirectElem.Type())
		newElem.Elem().Set(reflect.ValueOf(indirectElem.Interface()))
		fmt.Printf("[%s]-----------------------------------------\n", path)
		fmt.Printf("[%s]newElem: %v, type: %s, no: %d\n", path, newElem, newElem.Type(), newElem.Elem().NumField())
		elem := newElem.Elem()

		t.Run(elem.Type().Field(i).Name, func(t *testing.T) {
			fmt.Printf("[%s]field type: %v, kind: %s\n", path, elem.Field(i).Type(), elem.Field(i).Kind())
			fmt.Printf("[%s]struct field: %v, Tag: %s\n", path, elem.Type().Field(i), elem.Type().Field(i).Tag)
			if elem.Field(i).Kind() == reflect.Struct {
				fmt.Printf("[%s]Found struct dive\n", path)
				dive(t, path+"."+elem.Type().Field(i).Name, elem.Field(i), root)
			} else if elem.Field(i).Kind() == reflect.Slice {
				fmt.Printf("[%s]Found slice dive\n", path)
				for ii := 0; ii < elem.Field(i).Len(); ii++ {
					newField := reflect.New(elem.Field(i).Index(ii).Type())
					newField.Elem().Set(reflect.ValueOf(elem.Field(i).Index(ii).Interface()))
					dive(t, fmt.Sprintf("%s.%s.%d", path, elem.Type().Field(i).Name, ii), newField.Elem(), root)
				}
			} else if elem.Field(i).Kind() == reflect.String {
				fmt.Printf("[%s][Test] String\n", path)
				elem.Field(i).SetString("")
			} else if elem.Field(i).Kind() == reflect.Int {
				fmt.Printf("[%s][Test] Int\n", path)
				elem.Field(i).SetInt(0)
			} else {
				fmt.Printf("[%s][Test] Unknown\n", path)
				return
			}

			fmt.Printf("[%s]testElem: %v, type: %s\n", path, elem, newElem.Type())

			newRoot := reflect.New(root.Type())
			newRoot.Elem().Set(reflect.ValueOf(root.Interface()))
			cn := newRoot.Elem()
			currentNode := &cn
			nodes := strings.Split(path, ".")
			if len(nodes) == 1 {
				newRoot = elem
			} else {

				for index, node := range nodes {
					if node == "object" {
						continue
					}

					if index >= len(nodes)-2 {
						fmt.Printf("[%s]found at node: %s\n", path, node)
						if currentNode.FieldByName(node).Kind() == reflect.Slice {
							n, _ := strconv.Atoi(nodes[index+1])
							currentNode.FieldByName(node).Index(n).Set(elem)
						} else {
							currentNode.FieldByName(node).Set(elem)
						}
						break
					}

					fmt.Printf("[%s]currentNode: %v, type: %s at index %d/%d\n", path, currentNode, currentNode.Type(), index, len(nodes)-1)
					newCurrentNode := currentNode.FieldByName(node)
					currentNode = &newCurrentNode
				}
			}

			payload, err := json.Marshal(newRoot.Interface())
			if err != nil {
				panic(err)
			}
			fmt.Printf("[%s][Test] request payload: %s\n", path, string(payload))
		})
	}
}
