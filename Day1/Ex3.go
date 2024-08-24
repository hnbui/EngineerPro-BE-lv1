/// Viết chương trình nhập 1 slice số, in ra tổng, số lớn nhất, số nhỏ nhất, trung bình cộng, slice đã được sắp xếp (tăng/giảm)

package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	var s []int
	var n int
	var sum float64 = 0
	fmt.Print("Enter the number of elements in the slice: ")
	fmt.Scan(&n)
	fmt.Println("Enter elements of the slice: ")
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		s = append(s, x)
		sum += float64(x)
	}
	// sort.Ints(s) /// Ascending sort
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j] /// Descending sort: return s[i] < s[j] instead of s[i] > s[j]
	})
	fmt.Printf("Sum: %.3f\n", sum)
	fmt.Printf("Maximum: %d\n", slices.Max(s))
	fmt.Printf("Minimum: %d\n", slices.Min(s))
	fmt.Printf("Average: %.3f\n", sum/float64(len(s)))
	fmt.Println("Sorted slice:", s)
}
