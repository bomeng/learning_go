/*
1. Two Sum
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if index, prs := m[num]; prs {
			return []int{index, i}
		} else {
			m[target-num] = i
		}
	}
	return nil
}

func main() {
	array := []int{2, 7, 11, 15}
	const target = 9

	var result = twoSum(array, target)

	fmt.Print(result)
}
