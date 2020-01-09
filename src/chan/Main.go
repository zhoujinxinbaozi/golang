package main

import (
	"fmt"
)

func send(c chan int) {
	for i := 1; i < 10; i++ {
		c <- i
		fmt.Println("send data : ", i)
	}
}

func main() {
	resch := make(chan int, 20)
	strch := make(chan string, 10)
	send(resch)
	strch <- "wd"
	select {
	case a := <-resch:
		fmt.Println("get data Int : ", a)
	case b := <-strch:
		fmt.Println("get data String : ", b)
	default:
		fmt.Println("no channel actvie")

	}

}
