package visitor

import "fmt"

/*
The Visitor Design Pattern is a behavioral pattern that lets you
add new operations to existing object structures without modifying their classes.

It achieves this by allowing you to separate the algorithm from the objects it operates on.
*/

func TestVisitorPattern() {
	shapes := []Shape{
		NewCircle(5),
		NewRectangle(10, 4),
		NewCircle(2.5),
	}

	// Shape says:  "visitor, here I am — a Circle"
	// Visitor says: "great, I know what to do with a Circle"
	fmt.Println("=== Calculating Areas ===")
	areaCalculator := &AreaCalculatorVisitor{}
	for _, shape := range shapes {
		shape.Accept(areaCalculator)
	}

	fmt.Println("\n=== Exporting to SVG ===")
	svgExporter := &SvgExporterVisitor{}
	for _, shape := range shapes {
		shape.Accept(svgExporter)
	}
}
