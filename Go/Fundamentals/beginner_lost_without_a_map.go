package kata

func Maps(x []int) []int {
	var res []int
	for i := 0; i < len(x); i++ {
		res = append(res, x[i]*2)
	}
	return res
}
