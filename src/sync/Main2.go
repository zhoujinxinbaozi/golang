package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [10]int{4,5,6}
	fmt.Println(arr)
	fmt.Printf("add = %p", &arr)
	ff(&arr)
	fmt.Println(arr)
	fmt.Println(reflect.TypeOf(arr))
}

func ff(arr *[10]int){
	fmt.Printf("add = %p", arr)
	arr[0] = 100
}
