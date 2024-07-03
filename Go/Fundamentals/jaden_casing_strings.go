package kata

import "strings"

func ToJadenCase(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		if i == 0 && str[i] != ' ' {
			res += strings.ToUpper(string(str[i]))
		} else if str[i] == ' ' && str[i+1] != ' ' {
			res += string(str[i])
			res += strings.ToUpper(string(str[i+1]))
			i++
		} else {
			res += string(str[i])
		}
	}

	return res
}
