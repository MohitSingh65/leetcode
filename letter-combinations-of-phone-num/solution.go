package main

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	digit_to_char := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string

	var backtrack func(index int, current_combination string)
	backtrack = func(index int, current_combination string) {
		if index == len(digits) {
			result = append(result, current_combination)
			return
		}

		letters := digit_to_char[digits[index]]
		for letter := 0; letter < len(letters); letter++ {
			backtrack(index+1, current_combination+string(letters[letter]))
		}
	}
	backtrack(0, "")
	return result
}
