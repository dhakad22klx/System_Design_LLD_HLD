package chainofresponsibility

type FiftyDollarHandler struct {
	BaseCashHandler
}

func NewFiftyDollarHandler() *FiftyDollarHandler {
	return &FiftyDollarHandler{BaseCashHandler{denomination: 50}}
}
