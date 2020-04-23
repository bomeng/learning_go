/*
40. Combination Sum II

Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.
*/
package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	findCombination(candidates, target, 0, []int{}, &result)

	return result
}

func findCombination(candidates []int, target int, index int, list []int, result *[][]int) {
	if target == 0 {
		r := append([]int{}, list...)
		*result = append(*result, r)
		return
	}

	if target < 0 {
		return
	}

	for i := index; i < len(candidates); i++ {
		if i == index || candidates[i] != candidates[i-1] {
			list = append(list, candidates[i])
			findCombination(candidates, target-candidates[i], i+1, list, result)
			list = list[0 : len(list)-1]
		}
	}
}

func main() {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	fmt.Println(combinationSum2(candidates, target))
}
