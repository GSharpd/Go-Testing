package addresses

import "testing"

type testScenario struct {
	insertedAddress         string
	expectedReturnedAddress string
}

// Unit Tests
// Signature must begin with Test

func TestAddressType(t *testing.T) {
	testAddress := "Avenue whatever"

	expectedAddressType := "avenue"

	received := AddresType(testAddress)

	if received != expectedAddressType {
		t.Errorf("Type mismatch expected - expected %v got %v",
			expectedAddressType,
			received)
	}
}

func TestAddressTypes(t *testing.T) {
	scenarios := []testScenario{
		{"Street something", "street"},
		{"Plaza whatever", "plaza"},
		{"Avenue something", "avenue"},
		{"Road 123", "road"},
		{"Errado", "Address type is invalid"},
	}

	for _, scenario := range scenarios {

		addrType := AddresType(scenario.insertedAddress)
		if scenario.expectedReturnedAddress != addrType {
			t.Errorf("Type mismatch expected - expected %v got %v",
				scenario.expectedReturnedAddress,
				addrType)
		}
	}
}
