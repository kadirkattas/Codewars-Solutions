package kata

func DNAStrand(dna string) string {
	res := ""

	for _, nuc := range dna {
		switch nuc {
		case 'A':
			res += "T"
		case 'T':
			res += "A"
		case 'G':
			res += "C"
		case 'C':
			res += "G"
		}
	}
	return res
}
