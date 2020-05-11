package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	var result []int
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !m[i] {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	result := findDisappearedNumbers([]int{1, 2, 2})
	fmt.Println(result)
}
