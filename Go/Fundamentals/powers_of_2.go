package kata

func PowersOfTwo(n int) []uint64 {
	var res []uint64

	for power := 0; power <= n; power++ {
		if power == 0 {
			res = append(res, 1)
		} else {
			tempRes := 1
			for i := 1; i <= power; i++ {
				tempRes *= 2
			}
			res = append(res, uint64(tempRes))
		}
	}
	return res
}
