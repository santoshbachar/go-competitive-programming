package main

import (
	"fmt"
	"strconv"
)

/* https://www.codewars.com/kata/526571aae218b8ee490006f4
Bit Counting

Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number. You can guarantee that input is non-negative.

Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case
*/

func CountBits(num uint) int {
	n := int64(num)
	bits := strconv.FormatInt(n, 2)
	s := string(bits)
	fmt.Println("string for bits is ", s)
	sum := 0
	cs := "01"
	for i := 0; i < len(s); i++ {
		if s[i] == cs[1] {
			sum++
		}
	}
	fmt.Println(strconv.FormatInt(n, 2))
	return sum
}

func main() {
	fmt.Println("countBits()", CountBits(120253))
}
