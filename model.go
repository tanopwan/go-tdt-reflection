package main

// Shape ...
type Shape struct {
	Color string `test:"empty"`
	Edge  int    `test:"zero"`
}

// Foo ...
type Foo struct {
	FooString string
	Another   int
	Barfs     []Bar
}

// Bar ...
type Bar struct {
	Barf string
}

// MyShape ...
type MyShape struct {
	Foos  []Foo `test:"empty,dive"`
	Shape Shape `test:"dive"`
}
