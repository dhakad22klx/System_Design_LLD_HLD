//If the part makes no sense without the whole, use composition.

// Association is a general connection: two classes simply know about each other.
// Aggregation is a grouping: the whole and parts can exist independently.
// Composition is an ownership: the part’s existence is bound to the whole.

package main

import "fmt"

// LineItem represents a single product entry in an order
type LineItem struct {
	productName string
	quantity    int
	unitPrice   float64
}

func NewLineItem(productName string, quantity int, unitPrice float64) LineItem {
	return LineItem{productName: productName, quantity: quantity, unitPrice: unitPrice}
}

func (l LineItem) GetSubtotal() float64 {
	return float64(l.quantity) * l.unitPrice
}

func (l LineItem) GetProductName() string {
	return l.productName
}

func (l LineItem) Describe() {
	fmt.Printf("%s x%d @ $%.2f = $%.2f\n", l.productName, l.quantity, l.unitPrice, l.GetSubtotal())
}

// Order composes LineItem — it owns a slice of LineItems directly (not pointers)
type Order struct {
	orderId   string
	lineItems []LineItem
}

func NewOrder(orderId string) Order {
	return Order{orderId: orderId}
}

func (o *Order) AddItem(product string, quantity int, unitPrice float64) {
	o.lineItems = append(o.lineItems, NewLineItem(product, quantity, unitPrice))
}

func (o *Order) RemoveItem(product string) {
	filtered := o.lineItems[:0]
	for _, item := range o.lineItems {
		if item.GetProductName() != product {
			filtered = append(filtered, item)
		}
	}
	o.lineItems = filtered
}

func (o *Order) GetTotal() float64 {
	total := 0.0
	for _, item := range o.lineItems {
		total += item.GetSubtotal()
	}
	return total
}

func (o *Order) PrintReceipt() {
	fmt.Printf("Order: %s\n", o.orderId)
	for _, item := range o.lineItems {
		item.Describe()
	}
	fmt.Printf("Total: $%.2f\n", o.GetTotal())
}

func testCompsition() {
	order := NewOrder("ORD-1001")
	order.AddItem("Wireless Mouse", 2, 29.99)
	order.AddItem("USB-C Cable", 3, 9.99)
	order.AddItem("Laptop Stand", 1, 49.99)
	order.PrintReceipt()
	// Go's garbage collector handles cleanup automatically — no manual memory management needed.
}