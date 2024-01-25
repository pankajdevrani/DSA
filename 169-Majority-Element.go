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
