package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, exists := seen[nums[i]]; exists {
			return true
		}
		seen[nums[i]] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{3, 3})) // should print true
}
