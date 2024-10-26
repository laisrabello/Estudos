package main

func Rps(p1, p2 string) string {
	switch {
	case p1 == p2:
		return "Draw!"

	case (p1 == "scissors" && p2 == "paper" ||
		p1 == "paper" && p2 == "rock" ||
		p1 == "rock" && p2 == "scissors"):
		return "Player 1 won!"

	default:
		return "Player 2 won!"
	}
}
