/// Viêt hàm có input là 1 string, trả về map có key là kí tự, và value là số lần xuất hiện của kí tự đó trong string.

package main

import "fmt"

func countChar(s string) map[string]int {
	m := make(map[string]int)
	for _, char := range s {
		m[string(char)]++
	}
	return m
}

func main() {
	var s string
	fmt.Println("Enter a string: ")
	fmt.Scan(&s)
	fmt.Println(countChar(s))
}
