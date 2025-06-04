package main

func findWordsContaining(words []string, x byte) []int {
	empty := []int{}
	if len(words) == 0 {
		return empty
	}

	var result []int
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			if word[j] == x {
				result = append(result, i)
				break
			}
		}
	}

	return result
}
