package visitor

import "fmt"

// SVG export visitor — shapes untouched
type SvgExporterVisitor struct{}

func (v *SvgExporterVisitor) VisitCircle(circle *Circle) {
	fmt.Println("<circle r=\"" + fmt.Sprint(circle.GetRadius()) + "\" />")
}

func (v *SvgExporterVisitor) VisitRectangle(rectangle *Rectangle) {
	fmt.Println("<rect width=\"" + fmt.Sprint(rectangle.GetWidth()) +
		"\" height=\"" + fmt.Sprint(rectangle.GetHeight()) + "\" />")
}
