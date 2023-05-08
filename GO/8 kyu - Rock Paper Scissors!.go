// https://www.codewars.com/kata/5672a98bdbdd995fad00000f

package kata

func Rps(p1, p2 string) string {
	gameRules := [3][2]string{{"paper", "rock"}, {"rock", "scissors"}, {"scissors", "paper"}}
	for i := 0; i < len(gameRules); i++ {
		if p1 == gameRules[i][0] && p2 == gameRules[i][1] {
			return "Player 1 won!"
		}
		if p1 == gameRules[i][1] && p2 == gameRules[i][0] {
			return "Player 2 won!"
		}
	}
	return "Draw!"
}
