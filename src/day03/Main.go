package main

import (
	"fmt"
	"time"
)

func test(c chan int) {
	fmt.Println("test start")
	c <- 5
	fmt.Println("test end")
}

func testDeadLock(c chan int) {
	fmt.Println("testDeadLock start")
	fmt.Println(<-c)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("testDeadLock end")
}

//func main() {
////	//chanDemo()
////	c := make(chan int)
////	//go test(c)
////	//go testDeadLock(c)
////	//time.Sleep(time.Millisecond)
////	c <- 1
////	fmt.Println(123)
////}

func main() {
	ch := make (chan int, 1)  // 注意这里给的容量是1
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("通道channel已经满啦，塞不下东西了!")
	}
}
