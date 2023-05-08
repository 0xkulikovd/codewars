// https://www.codewars.com/kata/55a2d7ebe362935a210000b2

package kata

func SmallestIntegerFinder(numbers []int) int {
	length := len(numbers)
	num := numbers[0]
	for j := 1; j < length; j++ {
		if numbers[j] < num {
			num = numbers[j]
		}
	}
	return num
}
