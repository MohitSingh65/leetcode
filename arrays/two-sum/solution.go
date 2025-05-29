package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, found := seen[complement]; found {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}
