package main

import "fmt"

/* https://leetcode.com/problems/two-sum/
1. Two Sum
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
**/

func twoSumByBruteForce(nums []int, target int) []int {
	fmt.Println("Input ", nums, "target: ", target)
	arr := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			fmt.Println("i, v:", i, ",", nums[i], " j, v:", j, ",", nums[j], " total:", nums[i]+nums[j])
			if nums[i]+nums[j] == target {
				fmt.Println("found at i,j = ", i, ",", j)
				arr[0] = i
				arr[1] = j
				return arr
			}
		}
	}
	return arr
}

func twoSumByOnePassHashTable(nums []int, target int) []int {
	// pass 1
	// storing data in map
	m := make(map[int]int)
	// incorrect for input 3,2,4 and target 6
	// it takes first on 6-3 = 3 it takes 1st 3 which is incorrect
	/*for i, v := range nums {
		m[v] = i
	}*/
	fmt.Println(m)

	// second pass
	// evaluatiog using compliment
	// NOTE: range cannot be used as same element cannot be used twice
	// as per the question
	// wrong for input {3,3} target 6

	for i := 0; i < len(nums); i++ {
		if val, ok := m[target-nums[i]]; ok {
			fmt.Println("val=", val)
			return []int{i, val}
		}
		m[nums[i]] = i
	}
	/*for i, v := range m {
		if val, ok := m[target-i]; ok {
			return []int{v, val}
		}
	}*/

	return []int{-1, -1}
}

func twoSumByTwoPassHashTable(nums []int, target int) []int {
	var m map[int]int = make(map[int]int)
	for i, v := range m {
		fmt.Println("i,v", i, ",", v)
	}
	fmt.Println(len(m))
	for i := 0; i < len(nums); i = i + 1 {
		fmt.Println(len(m))
		var j int
		var ok bool
		j, ok = m[target-nums[i]]
		if ok {
			var ans []int = []int{i, j}
			return ans
		}
		m[nums[i]] = i
	}
	var ans []int = []int{-1, -1}
	return ans
}

func main() {
	input := []int{3, 2, 4}
	target := 40

	twoSumByBruteForce([]int{2, 5, 5, 11}, 10)
	//twoSum(input,target)
	fmt.Println("hash = ", twoSumByTwoPassHashTable(input, target))

	fmt.Println("onehash = ", twoSumByOnePassHashTable(input, target))

}
