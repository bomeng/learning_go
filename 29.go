/*
29. Divide Two Integers
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.
*/

package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}

	if dividend == 0 {
		return 0
	}

	ans := 0.0
	a := float64(dividend)
	b := float64(divisor)
	if dividend > 0 {
		if divisor > 0 {
			ans = math.Exp(math.Log(a) - math.Log(b))
		} else {
			ans = -math.Exp(math.Log(a) - math.Log(-b))
		}
	} else {
		if divisor > 0 {
			ans = -math.Exp(math.Log(-a) - math.Log(b))
		} else {
			ans = math.Exp(math.Log(-a) - math.Log(-b))
		}
	}

	return int(ans)
}

func main() {
	var result = divide(
		2147483647, 1)

	fmt.Print(result)
}
