package kata

func Invert(arr []int) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		res = append(res, arr[i]*-1)
	}
	return res
}
