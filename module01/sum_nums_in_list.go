package module01

func SumNumsInList(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + SumNumsInList(list[1:])
}
