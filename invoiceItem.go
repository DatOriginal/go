package main 

import (
	"fmt"
)

type InvoiceItem struct {
	id, desc string
	qty int64
	unitPrice float64
}

func NewInvoiceItem(id, desc string, qty int64, unitPrice float64) *InvoiceItem {
	return &InvoiceItem{id, desc, qty, unitPrice}
}

func (i *InvoiceItem) GetTotal() float64 {
	return float64(i.qty) * i.unitPrice
}

func main() {
	invi := NewInvoiceItem("A101", "Pen Red", 888, 0.08)
	invi = NewInvoiceItem("A101", "Pen Red", 999, 0.99)
	fmt.Printf(" Id: %s\n DESC: %s\n QTY: %v\n Unit Price: %.2f\n", invi.id, invi.desc, invi.qty, invi.unitPrice)
	fmt.Printf("The total: %.2f\n", invi.GetTotal())
}