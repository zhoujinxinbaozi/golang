package main

import (
	"fmt"
)
import "Entity"

func main() {
	//a1 := [...]int{6, 5, 4, 78, 3}
	//findMax(a1)
	//sliceTest()
	//var student1 entity.Student
	//student1.Age = 12
	//student1.Name = "zhoujinxin"
	//fmt.Println(student1)
	//i1 := 12
	//getAddress(&i1)
	//getStudent(&student1)
	mapTest()

}

func getAddress(i1 *int) {
	fmt.Println(*i1)
}

func getStudent(student *entity.Student) {
	fmt.Println(student.Name)
	fmt.Println(student.Age)
	fmt.Println(&student.Name)
}

func sliceTest() {
	s1 := make([]int, 0, 10)
	s1 = append(s1, 1)
	fmt.Print(s1)
	mapTest()
}

func mapTest() {
	m1 := make(map[string]int, 10)
	m1["123"] = 123
	m1["234"] = 234
	delete(m1, "123")
	//fmt.Print(m1)
	if index, ok := m1["234"]; ok {
		fmt.Println("contain")
		fmt.Println(index, ok) // index中存储的是m1["234"]的值
	} else {
   		fmt.Println("not contain")
	}
	value := true
	if value = false; value {
		fmt.Println(value)
	}else{
		fmt.Println("false")
	}

}

func findMax(array [5]int) {
	s1 := []int{5, 10}
	fmt.Println(s1)
	var tmp [5]int
	var max = -1
	var index = 0
	for _, i := range array {
		if i > max {
			tmp[index] = i
			index++
			max = i
		} else {
			tmp[index] = max
			index++
		}
	}
	fmt.Println(array)
	fmt.Println(tmp)
}
