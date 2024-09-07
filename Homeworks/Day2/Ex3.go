/// Viết hàm giải bài toán twosum: https://leetcode.com/problems/two-sum/ (sử dụng map)

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if _, ok := m[target-num]; ok {
			return []int{m[target-num], i}
		}
		m[num] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
