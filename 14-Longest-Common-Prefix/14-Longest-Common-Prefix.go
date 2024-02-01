// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	first := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(first) > len(strs[i]) {
			first = strs[i]
		}
	}
	println(first)
	result := ""
	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(first); j++ {
			if first[j] == strs[i][j] {
				result = result + string(first[j])

			} else {
				break
			}
		}
		first = result
		result = ""
	}
	return first
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
