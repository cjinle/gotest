package codewars

func TwoSum(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := len(numbers) - 1; j > i; j-- {
			if numbers[i]+numbers[j] == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
}
