package main

import (
	"fmt"
	"reflect"
)

var domains = []Domain{
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

func main() {

	type embeded struct {
		data interface{}
	}

	s := Shape{
		Color: "T",
		Edge:  1,
	}

	e := embeded{
		data: s,
	}
	v := reflect.ValueOf(&e.data)
	i := v.Elem()

	fmt.Printf("value of domain: %v, with type: %v\n", i, reflect.TypeOf(i.Interface()).Name())

	w := NewWalker()
	w.Init(domains)
}
