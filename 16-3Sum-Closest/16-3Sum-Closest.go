// Given an integer array nums of length n and an integer target,
// find three integers in nums such that the sum is closest to target.
// Return the sum of the three integers.
// Example 1:
// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

// Example 2:
// Input: nums = [0,0,0], target = 1
// Output: 0
// Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).

package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := 0
	for i := 0; i < len(nums)-2; i++ {
		result = nums[i] + nums[i+2] + nums[i+2]
		if result == target {
			return result
		} else if result < target {

		}
	}
	return result
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	target := 0
	fmt.Println("Majority Element:", threeSumClosest(nums, target))
}
