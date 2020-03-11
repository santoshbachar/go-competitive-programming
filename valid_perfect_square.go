package main

import "fmt"

/* https://leetcode.com/problems/valid-perfect-square/
367. Valid Perfect Square

Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:

Input: 16
Output: true
Example 2:

Input: 14
Output: false
*/

func isPerfectSquare(num int) bool {
	i, sum := 1, 0
	for i = 1; sum < num; i += 2 {
		sum += i
		if sum == num {
			return true
		}
	}
	return false
}

func main() {
	x := isPerfectSquare(36)

	/*for i := 0; i < 908202; i++ {
		if isPerfectSquare(i) {
			fmt.Println(i)
		}
	}*/

	fmt.Println("isPerfectSquare=", x)
}
