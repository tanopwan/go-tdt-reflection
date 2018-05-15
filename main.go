package main

// Shape ...
type Shape struct {
	Color string `test:"empty"`
	Edge  int    `test:"zero"`
}

// MyShape ...
type MyShape struct {
	Shapes []Shape `test:"empty,dive"`
	Shape  Shape   `test:"dive"`
}

func main() {

}
