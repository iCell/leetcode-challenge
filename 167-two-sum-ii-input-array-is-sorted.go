package main

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		switch {
		case sum == target:
			return []int{i + 1, j + 1}
		case sum > target:
			j--
		case sum < target:
			i++
		}
	}
	return nil
}
