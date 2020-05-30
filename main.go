package main

import (
	mdl01 "algorithms/module01"
	"fmt"
)

func main() {

	// Examples:
	fmt.Println(mdl01.DecimalToOtherBase(1, 2))   // "1"
	fmt.Println(mdl01.DecimalToOtherBase(2, 2))   // "10"
	fmt.Println(mdl01.DecimalToOtherBase(7, 3))   // "21"
	fmt.Println(mdl01.DecimalToOtherBase(14, 2))  // "1110"
	fmt.Println(mdl01.DecimalToOtherBase(14, 16)) // "E"
	fmt.Println(mdl01.DecimalToOtherBase(17, 16)) // "11"
}
