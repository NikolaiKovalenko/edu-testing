package sample

import "errors"

type Order struct {
	ID       string
	Quantity int
	Price    float64
}

type OrderProcessor interface {
	CalculateTotal(orders []Order) (float64, error)
}

type SimpleOrderProcessor struct{}

func (sop *SimpleOrderProcessor) CalculateTotal(orders []Order) (float64, error) {
	if len(orders) == 0 {
		return 0, errors.New("no orders provided")
	}

	var total float64
	for _, order := range orders {
		if order.Quantity < 0 || order.Price < 0 {
			return 0, errors.New("invalid order details")
		}
		total += float64(order.Quantity) * order.Price
	}

	return total, nil
}
