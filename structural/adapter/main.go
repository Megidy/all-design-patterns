package main

import "log"

func main() {

	adapter := NewPaymentAdapter(&PaymentSystem{})

	adapter.PayUSD(10)
}

type PaymentSystem struct{}

func NewPaymentSystem() *PaymentSystem { return &PaymentSystem{} }

func (ps *PaymentSystem) PayUAH(amount float64) {
	log.Printf("payed with UAH: %f", amount)
}

type PaymentProcessor interface {
	PayUSD(amount float64)
}

type PaymentAdapter struct {
	ps *PaymentSystem
}

func NewPaymentAdapter(ps *PaymentSystem) *PaymentAdapter {
	return &PaymentAdapter{
		ps: ps,
	}
}

func (pa *PaymentAdapter) PayUSD(amount float64) {
	pa.ps.PayUAH(amount * 40)
}
