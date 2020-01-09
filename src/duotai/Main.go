package main

import "fmt"

func main() {
	d := dog{"dog", "汪汪汪"}
	getCall(&d)

	c := cat{"cat", "喵喵喵"}
	getCall(&c)
}

func getCall(animal animal) {
	animal.call()
}

type animal interface {
	call()
}

type dog struct {
	name string
	yell string
}

type cat struct {
	name string
	yell string
}

func (cat *cat) call() {
	fmt.Println("===== cat call start =====")
	fmt.Printf("name is %s, yell is %s\n", cat.name, cat.yell)
	fmt.Println("===== cat call end =====")
}

func (dog *dog) call() {
	fmt.Println("===== dog call start =====")
	fmt.Printf("name is %s, yell is %s\n", dog.name, dog.yell)
	fmt.Println("===== dog call end =====")
}
