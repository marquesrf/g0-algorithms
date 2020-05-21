package main

import (
	mdl01 "algorithms/module01"
	"fmt"
)

func main() {

	// Examples:
	fmt.Println(mdl01.NumInList([]int{1, 2, 3, 4, 5}, 5))      // true
	fmt.Println(mdl01.NumInList([]int{3, 3, 3, 3, 3}, 5))      // false
	fmt.Println(mdl01.NumInList([]int{3, 3, 3, 5, 3}, 5))      // true
	fmt.Println(mdl01.NumInList([]int{4, 2, 22, -10, 8}, -10)) // true

	// Emplty lists!
	fmt.Println(mdl01.NumInList(nil, 5))     // false
	fmt.Println(mdl01.NumInList([]int{}, 5)) // false
}
