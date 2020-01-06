package main

import (
	"fmt"
)

func main() {
	//f1(f3(f2, 100, 200))
	//ff := f4()
	//ff(12)
	//ii := 12
	//fmt.Println(reflect.TypeOf(ff))
	//fmt.Println(reflect.TypeOf(ii))
	m1 := map[int]int{0:0, 1:1, 2:2}
	for k := range m1{
		fmt.Printf("key is %d, value is %d\n", k, m1[k])
	}
}

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Printf("x = %d, y = %d", x, y)
}

func f3(f func(int, int), x, y int) func() {
	return func() {
		f(x, y)
	}
}

func f4() func(int){
	return func(target int){
		fmt.Println(target)
	}
}


