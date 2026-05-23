package visitor

type ShapeVisitor interface {
	VisitCircle(circle *Circle)
	VisitRectangle(rectangle *Rectangle)
}
