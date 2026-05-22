package chainofresponsibility

// CashRequest carries the amount to dispense.
// Each handler reduces amount as it dispenses notes.
type CashRequest struct {
	amount int
}