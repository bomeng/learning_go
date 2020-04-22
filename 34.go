/*
34. Find First and Last Position of Element in Sorted Array

Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].
*/
package main

import "fmt"

func searchRange(nums []int, target int) []int {
	low := searchLow(nums, 0, len(nums)-1, target)
	if low == -1 {
		return []int{-1, -1}
	}
	high := searchHigh(nums, low, len(nums)-1, target)
	return []int{low, high}
}

func searchLow(nums []int, low int, high int, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if nums[mid] == target {
		if mid-1 >= 0 && nums[mid-1] == target {
			return searchLow(nums, low, mid-1, target)
		} else {
			return mid
		}
	}
	if target >= nums[mid] {
		return searchLow(nums, mid+1, high, target)
	} else {
		return searchLow(nums, low, mid-1, target)
	}
}

func searchHigh(nums []int, low int, high int, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if nums[mid] == target {
		if mid+1 <= high && nums[mid+1] == target {
			return searchHigh(nums, mid+1, high, target)
		} else {
			return mid
		}
	}
	if target >= nums[mid] {
		return searchHigh(nums, mid+1, high, target)
	} else {
		return searchHigh(nums, low, mid-1, target)
	}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	low := searchLow(nums, 0, len(nums)-1, target)
	high := searchHigh(nums, 0, len(nums)-1, target)
	fmt.Println(low, high)
}
