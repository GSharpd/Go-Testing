package addresses

import "strings"

// AddressType verifies if address is valid and returns it if true otherwise returns invalid type message
func AddresType(addr string) string {
	validTypes := []string{"street", "avenue", "road", "plaza"}

	firstWord := strings.ToLower(strings.Split(addr, " ")[0])

	isValid := false

	for _, validType := range validTypes {
		if firstWord == validType {
			isValid = true
		}
	}

	if isValid {
		return firstWord
	}

	return "Address type is invalid"
}
