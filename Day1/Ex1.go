/// Viết hàm nhập 2 cạnh của hình chữ nhật. In ra chu vi và diện tích.

package main

import (
	"fmt"
)

func main() {
	var w, h float64
	fmt.Print("Enter width: ")
	fmt.Scan(&w)
	fmt.Print("Enter height: ")
	fmt.Scan(&h)
	fmt.Printf("Perimeter: %.3f. Area: %.3f", 2*(w+h), (w * h))
}
