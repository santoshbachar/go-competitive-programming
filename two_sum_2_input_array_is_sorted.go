package main

import "fmt"

/* https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
167. Two Sum II - Input array is sorted

Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.
Example:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
*/

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{-1, -1}
}

func twoSumBetter(numbers []int, target int) []int {
	max := len(numbers) - 1
	min := 0
	sum := 0
	for {
		sum = numbers[min] + numbers[max]
		if sum == target {
			return []int{min + 1, max + 1}
		} else if sum > target {
			max--
		} else {
			min++
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println("twoSum=", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("twoSumBetter=", twoSumBetter([]int{2, 7, 11, 15}, 9))

}
