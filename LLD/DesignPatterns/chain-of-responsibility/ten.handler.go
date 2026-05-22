package chainofresponsibility

type TenDollarHandler struct {
	BaseCashHandler
}

func NewTenDollarHandler() *TenDollarHandler {
	return &TenDollarHandler{BaseCashHandler{denomination: 10}}
}