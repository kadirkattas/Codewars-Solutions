package kata

func Rps(p1, p2 string) string {
	switch p1 {
	case "rock":
		switch p2 {
		case "scissors":
			return "Player 1 won!"
		case "paper":
			return "Player 2 won!"
		}
	case "scissors":
		switch p2 {
		case "rock":
			return "Player 2 won!"
		case "paper":
			return "Player 1 won!"
		}
	case "paper":
		switch p2 {
		case "rock":
			return "Player 1 won!"
		case "scissors":
			return "Player 2 won!"
		}
	}
	return "Draw!"
}
