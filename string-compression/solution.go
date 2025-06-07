package main

import "fmt"

func compress(chars []byte) int {
	n := len(chars)
	write := 0
	i := 0

	for i < n {
		currentChar := chars[i]
		count := 0

		for i < n && chars[i] == currentChar {
			count++
			i++
		}

		chars[write] = currentChar
		write++

		if count > 1 {
			for _, digit := range []byte(fmt.Sprint(count)) {
				chars[write] = digit
				write++
			}
		}
	}
	return write
}
