/*
33. Search in Rotated Sorted Array

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).
*/
package main

import "fmt"

func search(nums []int, target int) int {
	return searchRecursive(nums, 0, len(nums)-1, target)
}

func searchRecursive(nums []int, low int, high int, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[low] <= nums[mid] {
		if target >= nums[low] && target <= nums[mid] {
			return searchRecursive(nums, low, mid-1, target)
		} else {
			return searchRecursive(nums, mid+1, high, target)
		}
	} else {
		if target >= nums[mid] && target <= nums[high] {
			return searchRecursive(nums, mid+1, high, target)
		} else {
			return searchRecursive(nums, low, mid-1, target)
		}
	}
}

func main() {
	arr := []int{3, 1}
	target := 1
	fmt.Print(search(arr, target))
}
