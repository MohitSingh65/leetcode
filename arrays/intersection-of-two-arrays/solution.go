package main

func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	set1 := make(map[int]struct{})
	seen := make(map[int]struct{})

	for _, num := range nums1 {
		set1[num] = struct{}{}
	}

	for _, num := range nums2 {
		if _, exists := set1[num]; exists {
			if _, alreadyAdded := seen[num]; !alreadyAdded {
				result = append(result, num)
				seen[num] = struct{}{}
			}
		}
	}
	return result
}
