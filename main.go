package main

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
	w := NewWalker()
	w.Init(domains)
}
