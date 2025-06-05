package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var empty []bool
	if len(candies) == 0 {
		return empty
	}
	max_candies := 0
	for _, c := range candies {
		if c > max_candies {
			max_candies = c
		}
	}

	result := make([]bool, len(candies))
	for i, c := range candies {
		result[i] = (c+extraCandies >= max_candies)
	}
	return result
}
