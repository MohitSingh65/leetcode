package main

func findDisappearedNumbers(nums []int) []int {
	numbers := make(map[int]struct{})
	for _, num := range nums {
		numbers[num] = struct{}{}
	}
	var missingNumbers []int

	for i := 1; i <= len(nums); i++ {
		if _, found := numbers[i]; !found {
			missingNumbers = append(missingNumbers, i)
		}
	}
	return missingNumbers
}
