package module01

func Reverse(someString string) string {
	var reverse string
	for _, r := range someString {
		reverse = string(r) + reverse
	}
	return reverse
}
