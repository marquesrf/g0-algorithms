package module01

func DecimalToOtherBase(dec, base int) string {
	var res string
	var charset = "0123456789ABCDEF"
	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec = dec / base
	}
	return res
}
