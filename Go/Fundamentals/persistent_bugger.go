package kata

func Persistence(n int) int {
	pers := 0

	for n >= 10 {
		mult := 1
		for n > 0 {
			rem := n % 10
			n = n / 10
			mult *= rem
		}
		n = mult
		pers++
	}

	return pers
}
