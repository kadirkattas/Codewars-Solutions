package kata

import "strconv"

func countSheep(num int) string {
	res := ""
	for n := 1; n <= num; n++ {
		res += strconv.Itoa(n) + " sheep..."
	}
	return res
}
