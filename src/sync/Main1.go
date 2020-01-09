package main

import (
	"fmt"
	"time"
)

func cal(a int, b int, exitChan chan int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	time.Sleep(time.Second * 2)
	exitChan <- c
}

func main() {

	exitChan := make(chan int, 8) //声明并分配管道内存
	for i := 0; i < 10; i++ {
		go cal(i, i+1, exitChan)
	}
	fmt.Println("==========")
	for j := 0; j < 10; j++ {
		fmt.Println(<-exitChan) //取信号数据，如果取不到则会阻塞
	}
	close(exitChan) // 关闭管道
}
