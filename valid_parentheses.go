package main

import (
	"strconv"
)

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

// Example 3:
// Input: s = "(]"
// Output: false

// Example 4:
// Input: s = "([])"
// Output: true

func StartValidParentheses() {
	s := "([}}])"

	println("Result: " + strconv.FormatBool(isValid(s)))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	parentheses := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	var orphans []string

	for i := 0; i < len(s); i++ {
		if i == 0 {
			for _, val := range parentheses {
				if val == string(s[i]) {
					return false
				}
			}
		}

		if _, found := parentheses[string(s[i])]; found {
			orphans = append(orphans, string(s[i]))
			continue
		}

		closeRef := ""

		for key, val := range parentheses {
			if val == string(s[i]) {
				closeRef = key
			}
		}

		if len(orphans) == 0 {
			return false
		}

		if closeRef == orphans[len(orphans)-1] {
			orphans = orphans[:len(orphans)-1]
		} else {
			return false
		}
	}

	return len(orphans) == 0
}
