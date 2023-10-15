package customer_test

import (
	"errors"
	"testing"

	"github.com/aletomasella/ddd-go/customer"
)

type testCase struct {
	testName    string
	name        string
	lastName    string
	expectedErr error
}

func TestCustomer_NewCustomer(t *testing.T) {

	newCustomerTestCases := []testCase{
		{
			testName:    "Valid name and lastname",
			name:        "Alejandro",
			lastName:    "Tomasella",
			expectedErr: nil,
		},
		{
			testName:    "Invalid name",
			name:        "",
			lastName:    "Tomasella",
			expectedErr: customer.ErrorInvalidName,
		},
		{
			testName:    "Invalid lastname",
			name:        "Alejandro",
			lastName:    "",
			expectedErr: customer.ErrorInvalidName,
		},
	}

	for _, tc := range newCustomerTestCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := customer.NewCustomer(tc.name, tc.lastName)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})

	}

}
