package module01

func SumNumsInList(list []int) int {
	var sum int
	for _, i := range list {
		sum += i
	}
	return sum
}
