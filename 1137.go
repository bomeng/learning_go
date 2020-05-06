package main

import (
	"fmt"
)

func tribonacci(n int) int {
	nums := make([]int, n + 2)
	nums[0] = 0
	nums[1] = 1
	nums[2] = 1

	if n <= 2 {
		return nums[3]
	}
	for i := 3; i <= n; i++ {
		nums[i] = nums[i - 1] + nums[i - 2] + nums[i - 3]
	}
	return nums[n]
}

func main() {
	fmt.Println(tribonacci(25))
}
