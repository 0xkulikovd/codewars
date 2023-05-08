// https://www.codewars.com/kata/5bb904724c47249b10000131

package kata

func Points(games []string) int {
	points := 0
	for _, str := range games {
		x := int(str[0])
		y := int(str[2])
		switch {
		case x > y:
			points += 3
		case x == y:
			points += 1
		}
	}
	return points
}
