package sample

import (
	"errors"
	"testing"
)

// MockPaymentProcessor is a mock implementation of the PaymentProcessor interface.
type MockPaymentProcessor struct {
	ProcessPaymentFn func(amount float64) (string, error)
}

func (mpp *MockPaymentProcessor) ProcessPayment(amount float64) (string, error) {
	return mpp.ProcessPaymentFn(amount)
}

func TestProcessPayment(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		amount          float64
		expectedOutcome string
		expectedError   error
		mockFn          func(amount float64) (string, error)
	}{
		{
			name:            "SuccessfulPayment",
			amount:          100.0,
			expectedOutcome: "mock_payment_id",
			expectedError:   nil,
			mockFn: func(amount float64) (string, error) {
				return "mock_payment_id", nil
			},
		},
		{
			name:            "FailedPaymentDueToZeroAmount",
			amount:          0.0,
			expectedOutcome: "",
			expectedError:   errors.New("invalid amount"),
			mockFn: func(amount float64) (string, error) {
				return "", errors.New("invalid amount")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProcessor := &MockPaymentProcessor{
				ProcessPaymentFn: tt.mockFn,
			}

			outcome, err := mockProcessor.ProcessPayment(tt.amount)

			if outcome != tt.expectedOutcome {
				t.Errorf("expected outcome %s, got %s", tt.expectedOutcome, outcome)
			}

			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}

// OUTPUT:
//=== RUN   TestProcessPayment
//=== RUN   TestProcessPayment/SuccessfulPayment
//=== RUN   TestProcessPayment/FailedPaymentDueToZeroAmount
//--- PASS: TestProcessPayment (0.00s)
//--- PASS: TestProcessPayment/SuccessfulPayment (0.00s)
//--- PASS: TestProcessPayment/FailedPaymentDueToZeroAmount (0.00s)
//PASS
//
//Process finished with the exit code 0
