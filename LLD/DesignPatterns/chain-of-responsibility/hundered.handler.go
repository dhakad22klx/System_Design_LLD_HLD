package chainofresponsibility

type HundredDollarHandler struct {
	BaseCashHandler
}

func NewHundredDollarHandler() *HundredDollarHandler {
	return &HundredDollarHandler{BaseCashHandler{denomination: 100}}
}
