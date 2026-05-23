package visitor

import (
	"fmt"
	"math"
)

// Area visitor — knows how to calculate area for each shape
type AreaCalculatorVisitor struct{}

func (v *AreaCalculatorVisitor) VisitCircle(circle *Circle) {
	area := math.Pi * circle.GetRadius() * circle.GetRadius()
	fmt.Println("Area of Circle: " + fmt.Sprint(area))
}

func (v *AreaCalculatorVisitor) VisitRectangle(rectangle *Rectangle) {
	area := rectangle.GetWidth() * rectangle.GetHeight()
	fmt.Println("Area of Rectangle: " + fmt.Sprint(area))
}
