package main

import "fmt"

/* https://leetcode.com/problems/sqrtx/
69. Sqrt(x)

Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
			 the decimal part is truncated, 2 is returned.
*/

func mySqrt(x int) int {

	if x <= 0 {
		return 0
	}
	if x == 1 || x == 2 {
		return 1
	}

	i := 2
	sqr := i * i

	for {

		if sqr == x {
			return i
		} else if sqr > x {
			fmt.Println("sqr > x so return i-1", sqr, x, i, i-1)
			return i - 1
		}
		i++
		sqr = i * i
	}
}

func main() {
	fmt.Println("sqrtx=", mySqrt(8))
}
