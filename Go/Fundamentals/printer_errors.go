package kata

import "strconv"

func PrinterError(s string) string {
	res := "/" + strconv.Itoa(len(s))
	errCount := 0
	for _, char := range s {
		if char > 'm' {
			errCount++
		}
	}
	res = strconv.Itoa(errCount) + res
	return res
}
