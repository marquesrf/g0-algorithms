package module01

func Reverse(someString string) string {
	var reverse string
	for i := len(someString) - 1; i >= 0; i-- {
		reverse = reverse + string(someString[i])
	}
	return reverse
}
