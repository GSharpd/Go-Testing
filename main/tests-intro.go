package main

import (
	"fmt"
	"modulo/addresses"
)

func main() {
	addrType := addresses.AddresType("Road whatever")

	fmt.Println(addrType)

	addrType2 := addresses.AddresType("Street whatever")

	fmt.Println(addrType2)

	addrType3 := addresses.AddresType("avenue whatever")

	fmt.Println(addrType3)

	addrType4 := addresses.AddresType("Rua whatever")

	fmt.Println(addrType4)
}
