package kata

func GetSum(a, b int) int {
	res := 0

	var starter int
	var finisher int

	if a > b {
		starter = b
		finisher = a
	} else {
		starter = a
		finisher = b
	}

	for starter <= finisher {
		res += starter
		starter++
	}

	return res
}
