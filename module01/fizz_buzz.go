package module01

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i < n; i++ {
		printValue(i)
		fmt.Print(", ")
	}
	printValue(n)
	fmt.Println()
}

func printValue(n int) {
	switch {
	case n%3 == 0 && n%5 == 0:
		fmt.Print("Fizz Buzz")
	case n%3 == 0:
		fmt.Print("Fizz")
	case n%5 == 0:
		fmt.Print("Buzz")
	default:
		fmt.Print(n)
	}
}
