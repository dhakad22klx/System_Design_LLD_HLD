package chainofresponsibility

type TwentyDollarHandler struct {
	BaseCashHandler
}

func NewTwentyDollarHandler() *TwentyDollarHandler {
	return &TwentyDollarHandler{BaseCashHandler{denomination: 20}}
}
