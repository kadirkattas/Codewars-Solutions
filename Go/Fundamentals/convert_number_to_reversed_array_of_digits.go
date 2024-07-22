package kata

func Digitize(n int) []int {
	var res []int

	for n > 0 {
		rem := n % 10
		res = append(res, rem)
		n = n / 10
	}

	if res != nil {
		return res
	}
	return []int{0}
}
