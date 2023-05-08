// https://www.codewars.com/kata/56fa3c5ce4d45d2a52001b3c

package kata

func Xor(a, b bool) bool {
	if (!a && b) || (a && !b) {
		return true
	}
	return false
}
