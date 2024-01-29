// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
// Example 1:
// Input: nums = [3,2,3]
// Output: 3
// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2
package main

import "fmt"

func majorityElement(nums []int) int {
	majorityElement := nums[0]
	count := 1

	for _, num := range nums[1:] {
		if count == 0 {
			majorityElement = num
		}
		if num == majorityElement {
			count++
		} else {
			count--
		}
	}

	return majorityElement
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("Majority Element:", majorityElement(nums))
}
