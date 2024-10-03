package main

import (
	"strconv"
)

// Case 1
// input := "abcabcbb"
// output 3

// Case 2
// input: s = "bbbbb"
// output: 1

// Case 3
// input: s = "pwwkew"
// output: 3

// Case 4 <-- Really -.-
// input: s = " "
// output: 1

// Case 5 <-- Can't believe it...
// input: s = " "
// output: 1

// Case 5 <-- Okay, my bad
// input: s = "au"
// output: 2

// Case 6 <-- STOP!
// input: s = "dvdf"
// output: 3

func StartLongestSubstring() {
	input := "dvdf"

	response := lengthOfLongestSubstring(input)

	println("The longest substring is: " + strconv.Itoa(response))
}

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	longest := 0
	start := 0

	for i, char := range s {

		if lastIndex, found := charMap[char]; found && lastIndex >= start {
			start = lastIndex + 1
		}

		charMap[char] = i

		currentLength := i - start + 1
		if currentLength > longest {
			longest = currentLength
		}
	}

	return longest
}
