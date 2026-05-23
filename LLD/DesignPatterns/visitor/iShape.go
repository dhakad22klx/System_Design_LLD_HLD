package visitor

type Shape interface {
	Accept(visitor ShapeVisitor) // "here I am, do what you(visitor) want"
}
