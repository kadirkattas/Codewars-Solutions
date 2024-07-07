package kata

func TwoToOne(s1 string, s2 string) string {
	combStr := s1 + s2
	res := ""

	for _, char := range combStr {
		alreadyAdded := false
		for _, letter := range res {
			if letter == char {
				alreadyAdded = true
				break
			}
		}
		if !alreadyAdded {
			res += string(char)
		}
	}

	res = SortLetters(res)
	return res
}

func SortLetters(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}
