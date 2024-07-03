package kata

func IsValidWalk(walk []rune) bool {
	locationX := 0
	locationY := 0

	for step := 0; step < len(walk); step++ {
		switch walk[step] {
		case 'w':
			locationX--
		case 'e':
			locationX++
		case 'n':
			locationY++
		case 's':
			locationY--
		}
	}

	return locationX == 0 && locationY == 0 && len(walk) == 10

	return false
}
