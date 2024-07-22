package kata

func StringToArray(str string) []string {
	var res []string
	tempWord := ""
	for _, char := range str {
		if char != ' ' {
			tempWord += string(char)
		} else {
			res = append(res, tempWord)
			tempWord = ""
		}
	}
	res = append(res, tempWord)
	return res
}
