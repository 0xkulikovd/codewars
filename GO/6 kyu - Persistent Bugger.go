// https://www.codewars.com/kata/55bf01e5a717a0d57e0000ec

package kata

import "strconv"

func MultiplyDigits(num int) int {
	arr := strconv.Itoa(num)
	res := 1
	var n int
	for _, v := range arr {
		n, _ = strconv.Atoi(string(v))
		res *= n
	}
	return res
}

func Persistence(n int) int {
	pers := 0

	for i := true; i; {
		if len(strconv.Itoa(n)) == 1 {
			i = false
		} else {
			n = MultiplyDigits(n)
			pers += 1
		}
	}

	return pers
}
