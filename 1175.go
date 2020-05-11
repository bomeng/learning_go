package main

import "fmt"

func numPrimeArrangements(n int) int {
	var primes []int
	nums := make([]bool, n + 1)
	for i := 2; i <= n; i++ {
		if nums[i] {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k <= n; k += i {
			nums[k] = true
		}
	}
	i := len(primes)
	j := n - i
	sum := 1
	for k := i; k >= 1; k-- {
		sum = sum * k
		sum %= 1000000007
	}
	for k := j; k >= 1; k-- {
		sum = sum * k
		sum %= 1000000007
	}

	return sum
}

func main() {
	i := numPrimeArrangements(100)
	fmt.Println(i)
}
