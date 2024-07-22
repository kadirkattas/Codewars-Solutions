package kata

func Grow(arr []int) int {
	res := 1
	for _, num := range arr {
		res *= num
	}
	return res
}
