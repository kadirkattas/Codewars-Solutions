package kata

func NearestSq(n int) int {
	nextSq := FindNextSq(n)
	prevSq := FindPrevSq(n)

	if n-prevSq < nextSq-n {
		return prevSq
	}

	return nextSq
}

func FindNextSq(n int) int {
	for i := 1; i < n; i++ {
		if i*i > n {
			return i * i
		}
	}
	return 1
}

func FindPrevSq(n int) int {
	for i := n; i > 0; i-- {
		if i*i < n {
			return i * i
		}
	}
	return 1
}
