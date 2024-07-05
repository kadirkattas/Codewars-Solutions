package kata

func Summation(n int) int {
	res := 0
	for n > 0 {
		res += n
		n = n - 1
	}
	return res
}
