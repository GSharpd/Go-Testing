package addresses

import "testing"

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
