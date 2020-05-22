package module01

import (
	"strings"
)

func Reverse(someString string) string {
	var sb strings.Builder
	for i := len(someString) - 1; i >= 0; i-- {
		sb.WriteByte(someString[i])
	}
	return sb.String()
}
