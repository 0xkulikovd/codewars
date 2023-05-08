// https://www.codewars.com/kata/5aa736a455f906981800360d

package kata

func Feast(beast string, dish string) bool {
	if (beast[0] == dish[0]) && (beast[len(beast)-1] == dish[len(dish)-1]) {
		return true
	}
	return false
}
