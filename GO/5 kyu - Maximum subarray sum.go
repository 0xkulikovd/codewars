// https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c

package kata

func SubarraySum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func MaximumSubarraySum(numbers []int) int {
	max := 0
	var sum int
	for i, v := range numbers {
		if v <= 0 {
			continue
		}
		for j := i + 1; j <= len(numbers); j++ {
			sum = SubarraySum(numbers[i:j])
			if sum <= 0 {
				break
			}
			if sum > max {
				max = sum
			}
		}
	}

	return max
}
