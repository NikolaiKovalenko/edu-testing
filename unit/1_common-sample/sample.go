package sample

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64) (string, error)
}

type RealPaymentProcessor struct{}

func (rpp *RealPaymentProcessor) ProcessPayment(amount float64) (string, error) {
	if amount <= 0 {
		return "", fmt.Errorf("invalid amount: %f", amount)
	}
	// Simulate a successful payment processing.
	return "payment_id_12345", nil
}
