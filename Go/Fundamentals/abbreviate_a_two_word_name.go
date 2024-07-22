package kata

import "strings"

func AbbrevName(name string) string {
	res := ""

	for i := 0; i < len(name); i++ {
		if (i == 0 && name[i] != ' ') || name[i-1] == ' ' {
			res += strings.ToUpper(string(name[i])) + "."
		}
	}

	return res[:len(res)-1]
}
