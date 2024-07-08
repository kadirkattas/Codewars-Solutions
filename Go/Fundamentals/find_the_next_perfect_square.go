package kata

func FindNextSquare(sq int64) int64 {
	number := int(sq)
	root := -1

	for num := 1; num*num <= number; num++ {
		if num*num == number {
			root = num
			break
		}
	}

	if root == -1 {
		return int64(root)
	}

	nextRoot := root + 1
	return int64(nextRoot * nextRoot)
}
