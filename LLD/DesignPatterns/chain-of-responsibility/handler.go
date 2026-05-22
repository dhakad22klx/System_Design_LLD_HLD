package chainofresponsibility

import "fmt"

// CashHandler is the contract every node in the chain must satisfy.
type CashHandler interface {
	setNext(next CashHandler)
	dispense(request *CashRequest)
}

// BaseCashHandler provides default chaining and dispensing logic.
// Concrete handlers embed this and set their denomination.
type BaseCashHandler struct {
	next         CashHandler
	denomination int
}

func (h *BaseCashHandler) setNext(next CashHandler) {
	h.next = next
}

func (h *BaseCashHandler) dispense(request *CashRequest) {
	if request.amount >= h.denomination {
		noteCount := request.amount / h.denomination
		request.amount = request.amount % h.denomination
		fmt.Printf("Dispensing %d x $%d\n", noteCount, h.denomination)
	}
	h.forward(request)
}

// forward passes the request to the next handler if one exists.
func (h *BaseCashHandler) forward(request *CashRequest) {
	if h.next != nil {
		h.next.dispense(request)
	}
}
