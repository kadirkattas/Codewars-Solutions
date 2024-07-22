package kata

func ReverseSeq(n int) []int {
	res := make([]int, n)

	for i := n; i > 0; i-- {
		res[n-i] = i
	}
	return res
}
