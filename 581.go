package main

import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	i := 0
	for i < len(nums) {
		if nums[i] != sorted[i] {
			break
		}
		i++
	}

	j := len(nums) - 1
	for j >=i {
		if nums[j] != sorted[j] {
			break
		}
		j--
	}

	return j - i + 1
}

func main() {
	i := findUnsortedSubarray([]int {1, 2})
	fmt.Println(i)
}