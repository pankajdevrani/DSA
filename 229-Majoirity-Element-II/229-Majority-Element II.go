// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
// Example 1:
// Input: nums = [3,2,3]
// Output: [3]
// Example 2:
// Input: nums = [1,2]
// Output: [1,2]
package main

import "fmt"

func majorityElement2(nums []int) []int {

	return []int{0, 2}
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("Majority Element:", majorityElement2(nums))
}
