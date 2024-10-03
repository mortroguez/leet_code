package main

// Example 1:

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]
// Example 2:

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]
// Explanation: [9,4] is also accepted.

func StartArrayIntersection() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	intersect(nums1, nums2)
}

func intersect(nums1 []int, nums2 []int) []int {
	numbersMap := make(map[int]int)
	var intersection []int

	for i := 0; i < len(nums1); i++ {
		numbersMap[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		if count, found := numbersMap[nums2[i]]; found && count > 0 {
			intersection = append(intersection, nums2[i])
			numbersMap[nums2[i]]--
		}
	}

	return intersection
}
