// https://www.codewars.com/kata/5a34b80155519e1a00000009

package kata

func multipleOfIndex(ints []int) []int {
	arr := []int{}
	for i, v := range ints {
		if i != 0 && v%i == 0 {
			arr = append(arr, v)
		}
	}
	return arr
}
