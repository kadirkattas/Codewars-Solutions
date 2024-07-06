package kata

func Order(sentence string) string {
	var words []string
	var tempWord string

	for _, char := range sentence {
		if char == ' ' {
			words = append(words, tempWord)
			tempWord = ""
		} else {
			tempWord += string(char)
		}
	}

	if tempWord != "" {
		words = append(words, tempWord)
	}

	for !CheckSort(words) {
		words = SortTheWords(words)
	}

	// 4 1 6 3 5 2 --> start
	// 3 1 6 4 5 2
	// 1 3 6 4 5 2
	// 1 3 2 4 5 6
	// 1 3 2 4 5 6
	// 1 3 2 4 5 6
	// 1 3 2 4 5 6 --> must return to back for 3

	res := ""

	for index, word := range words {
		if index == len(words)-1 {
			res += word
		} else {
			res += word + " "
		}
	}

	return res
}

func SortTheWords(words []string) []string {
	for i := 0; i < len(words); i++ {
		for _, char := range words[i] {
			if char > '0' && char <= '9' {
				words[int(char-'0')-1], words[i] = words[i], words[int(char-'0')-1]
				break
			}
		}
	}
	return words
}

func CheckSort(words []string) bool {
	for i := 0; i < len(words); i++ {
		for _, char := range words[i] {
			if char > '0' && char <= '9' {
				if int(char-'0')-1 != i {
					return false
				}
			}
		}
	}
	return true
}
