package kata

func DNAtoRNA(dna string) string {
	res := ""
	for _, nAcide := range dna {
		if nAcide == 'T' {
			res += "U"
		} else {
			res += string(nAcide)
		}
	}
	return res
}
