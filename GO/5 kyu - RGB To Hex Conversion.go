// https://www.codewars.com/kata/513e08acc600c94f01000001

package kata

import "strconv"

func RGB(r, g, b int) string {

	hex := toHex(r) + toHex(g) + toHex(b)

	return hex
}

func toHex(num int) string {

	if num <= 0 {
		return "00"
	}
	if num > 255 {
		return "FF"
	}

	d := map[int]string{10: "A", 11: "B", 12: "C", 13: "D", 14: "E", 15: "F"}

	second := num % 16
	first := (num - second) / 16

	s := ""

	if first < 10 {
		s = s + strconv.Itoa(first)
	} else {
		s = s + d[first]
	}

	if second < 10 {
		s = s + strconv.Itoa(second)
	} else {
		s = s + d[second]
	}

	return s
}
