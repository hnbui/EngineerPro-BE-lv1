/// Viết chương trình giải bài toán twosum: https://leetcode.com/problems/two-sum/

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums { /// Other syntax: for i := 0; i < len(nums); i++ {num := nums[i]}
		complement := target - num
		if index, ok := numMap[complement]; ok { /// Other syntax: index, ok := numMap[complement] // if ok {}
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	var nums []int
	var target int
	fmt.Print("Enter the list of numbers (space-separated): ")
	fmt.Scan(&nums)
	fmt.Print("Enter the target number: ")
	fmt.Scan(&target)
	result := twoSum(nums, target)
	if result == nil {
		fmt.Println("No pair of numbers sum up to the target.")
	} else {
		fmt.Printf("The pair of numbers that sum up to the target are at indices %d and %d.\n", result[0], result[1])
	}
}
