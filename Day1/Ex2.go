/// Viết chương trình nhập 1 string, in ra true nếu độ dài chuỗi chia hết cho 2, và false nếu ngược lại.

package main

import (
	"fmt"
)

func main() {
	var inputString string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&inputString)

	if len(inputString)%2 == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
