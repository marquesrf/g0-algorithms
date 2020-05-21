package module01

// Return true if  num is in the list slice, false otherwise.
func NumInList(list []int, num int) bool {
	for _, i := range list {
		if i == num {
			return true
		}
	}
	return false
}
