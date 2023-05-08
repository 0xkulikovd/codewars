// https://www.codewars.com/kata/5727bb0fe81185ae62000ae3

package kata

func CleanString(s string) string {
	acc := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "#" {
			if acc == "" {
				continue
			}
			acc = acc[:len(acc)-1]
		} else {
			acc += string(s[i])
		}
	}
	return acc
}
