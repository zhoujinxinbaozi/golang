package main

import (
	"fmt"
)

// golang 中map和slice是引用传递
func main() {
	slice := []int{1, 2, 3}
	fmt.Printf("address = %p\n", &slice)
	testSlice(slice)
	fmt.Println(slice)
}

func testSlice(slice []int) {
	fmt.Printf("address = %p\n", &slice)
	slice[0] = 2
	fmt.Println(slice)
}
