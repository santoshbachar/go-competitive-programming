package main

import (
	"fmt"
	"math"
)

/* https://leetcode.com/problems/sum-of-square-numbers/
633. Sum of Square Numbers

Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a2 + b2 = c.

Example 1:

Input: 5
Output: True
Explanation: 1 * 1 + 2 * 2 = 5


Example 2:

Input: 3
Output: False
*/

func judgeSquareSum(c int) bool {
	if c == 0 {
		return true
	}

	var high, low int = int(math.Sqrt(float64(c))), 0
	fmt.Println(high, low)

	sum := high*high + low*low

	for {
		if low > high {
			break
		}

		fmt.Println("high,low", high, low)
		if sum == c {
			return true
		} else if sum < c {
			low++
		} else {
			high--
		}
		sum = high*high + low*low
	}
	return false
}

func main() {
	fmt.Println("judgeSquareSum=", judgeSquareSum(1))
}
