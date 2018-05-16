# go-tdt-reflection

### Keywords
- Golang
- Table Driven Tests
- Testing
- Reflection through struct

### Example
```
// Struct
type Shape struct {
	Color string `test:"empty"`
	Edge  int    `test:"zero"`
}

type MyShape struct {
	Shapes []Shape `test:"empty,dive"`
	Shape  Shape   `test:"dive"`
}
```

```
// Struct to reflect
MyShape{
  Shapes: []Shape{
    Shape{Color: "Blue", Edge: 3},
    Shape{Color: "Purple", Edge: 2}
  },
  Shape: Shape{Color: "Pink", Edge: 11},
}
```

```
--- PASS: TestDomains (0.00s)
    --- PASS: TestDomains/MyShape (0.00s)
        --- PASS: TestDomains/MyShape/Shapes (0.00s)
            --- PASS: TestDomains/MyShape/Shapes/Color (0.00s)
            --- PASS: TestDomains/MyShape/Shapes/Edge (0.00s)
            --- PASS: TestDomains/MyShape/Shapes/Color#01 (0.00s)
            --- PASS: TestDomains/MyShape/Shapes/Edge#01 (0.00s)
        --- PASS: TestDomains/MyShape/Shape (0.00s)
            --- PASS: TestDomains/MyShape/Shape/Color (0.00s)
            --- PASS: TestDomains/MyShape/Shape/Edge (0.00s)
```

Each test will generate json request body with all possible cases specify in tags
```
// TestDomains/MyShape/Shapes/Color
// [Test] request payload:
{
  "Shapes":[
    {
      "Color":"", // tag: empty
      "Edge":3
    },
    {
      "Color":"Purple",
      "Edge":2
    }
  ],
  "Shape":{
    "Color":"Pink"
    ,"Edge":11
  }
}
```