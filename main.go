package main

import (
	mdl01 "algorithms/module01"
	"fmt"
)

func main() {

	// Examples:
	fmt.Println(mdl01.SumNumsInList([]int{1, 2, 3, 4, 5}))    // 15
	fmt.Println(mdl01.SumNumsInList([]int{3, 3, 3, 3, 3}))    // 15
	fmt.Println(mdl01.SumNumsInList([]int{3, 5, 3, 5, 3}))    // 19
	fmt.Println(mdl01.SumNumsInList([]int{4, 2, 22, -10, 8})) // 26

	// Emplty lists!
	fmt.Println(mdl01.SumNumsInList(nil))     // 0
	fmt.Println(mdl01.SumNumsInList([]int{})) // 0
}
