package kata

func Solution(str string) []string {
	var res []string

	for i := 0; i < len(str)-1; i++ {
		res = append(res, string(str[i])+string(str[i+1]))
		i++
	}

	if len(str)%2 != 0 {
		res = append(res, string(str[len(str)-1])+"_")
	}

	return res
}
