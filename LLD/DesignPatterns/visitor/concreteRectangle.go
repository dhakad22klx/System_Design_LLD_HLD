package visitor

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{width: width, height: height}
}

func (r *Rectangle) GetWidth() float64 {
	return r.width
}

func (r *Rectangle) GetHeight() float64 {
	return r.height
}

func (r *Rectangle) Accept(visitor ShapeVisitor) {
	visitor.VisitRectangle(r) // "I am a Rectangle — here I am"
}
