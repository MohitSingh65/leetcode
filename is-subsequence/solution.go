package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	if s[0] == t[0] {
		return isSubsequence(s[1:], t[1:])
	}

	return isSubsequence(s, t[1:])
}

