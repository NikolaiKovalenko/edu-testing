package sample

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTotal(t *testing.T) {
	t.Parallel()
	processor := &SimpleOrderProcessor{}

	tests := []struct {
		name          string
		orders        []Order
		expectedTotal float64
		expectedError string
	}{
		{
			name: "ValidOrders",
			orders: []Order{
				{ID: "1", Quantity: 2, Price: 50.0},
				{ID: "2", Quantity: 1, Price: 100.0},
			},
			expectedTotal: 200.0,
			expectedError: "",
		},
		{
			name:          "NoOrders",
			orders:        nil,
			expectedTotal: 0.0,
			expectedError: "no orders provided",
		},
		{
			name: "InvalidOrderDetails",
			orders: []Order{
				{ID: "1", Quantity: -1, Price: 50.0},
			},
			expectedTotal: 0.0,
			expectedError: "invalid order details",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total, err := processor.CalculateTotal(tt.orders)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError, "error message mismatch")
			} else {
				assert.NoError(t, err, "unexpected error occurred")
				assert.Equal(t, tt.expectedTotal, total, "total price mismatch")
			}
		})
	}
}

//OUTPUT:
//=== RUN   TestCalculateTotal
//=== RUN   TestCalculateTotal/ValidOrders
//=== RUN   TestCalculateTotal/NoOrders
//=== RUN   TestCalculateTotal/InvalidOrderDetails
//--- PASS: TestCalculateTotal (0.00s)
//--- PASS: TestCalculateTotal/ValidOrders (0.00s)
//--- PASS: TestCalculateTotal/NoOrders (0.00s)
//--- PASS: TestCalculateTotal/InvalidOrderDetails (0.00s)
//PASS
//
//Process finished with the exit code 0
