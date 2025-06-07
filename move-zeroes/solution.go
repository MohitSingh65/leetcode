package main

func moveZeroes(nums []int) {
	if nums == nil {
		return
	}

	writeIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[writeIndex] = nums[i]
			if i != writeIndex {
				nums[i] = 0
			}
			writeIndex++
		}
	}
	return
}
