package kata

import "strings"

func ToCamelCase(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		if s[i] == '-' || s[i] == '_' {
			res += strings.ToUpper(string(s[i+1]))
			i++
		} else {
			res += string(s[i])
		}
	}

	return res
}
