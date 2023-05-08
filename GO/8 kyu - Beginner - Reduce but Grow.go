// https://www.codewars.com/kata/57f780909f7e8e3183000078

package kata

func Grow(arr []int) int {
	result := 1
	for _, num := range arr {
		result *= num
	}
	return result
}
