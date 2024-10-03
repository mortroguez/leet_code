package main

import "strconv"

// Case 1
// nums := []int{2,7,11,15}
// target := 9
// output [0,1]

// Case 2
// nums := []int{3,2,4}
// target := 6
// output [1,2]

// Case 3
// nums := []int{3,3}
// target := 6
// output [0,1]

// Case 4
// nums := []int{3,2,3}
// target := 6
// output [0,2]

func startTwoSum() {
	// 1 - Two sums
	nums := []int{3, 2, 3}
	target := 6
	response := twoSum(nums, target)

	for i := 0; i < len(response); i++ {
		print(strconv.Itoa(response[i]) + " ")
	}
}

func twoSum(nums []int, target int) []int {
	var response []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				response = append(response, i)
				response = append(response, j)
			}
		}
	}

	return response
}
