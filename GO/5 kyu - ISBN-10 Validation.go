// https://www.codewars.com/kata/51fc12de24a9d8cb0e000001

package kata

import (
	"regexp"
	"strconv"
)

func ValidISBN10(isbn string) bool {
	if len(isbn) < 10 {
		return false
	}
	acc := 0
	for i := 0; i < len(isbn); i++ {

		var num int

		match, _ := regexp.MatchString("([0-9])", string(isbn[i]))
		matchX, _ := regexp.MatchString("([0-9]|X|x)", string(isbn[i]))

		if (i < 9 && !(match)) || (i == 9 && !(matchX)) {
			return false
		}

		if isbn[i] == 'X' || isbn[i] == 'x' {
			num = 10
		} else {
			num, _ = strconv.Atoi(string(isbn[i]))
		}

		acc += num * (i + 1)
	}
	return acc%11 == 0
}
