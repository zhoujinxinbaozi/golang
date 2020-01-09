package main

import (
	"Entity"
	"fmt"
	"reflect"
)

func main() {
	i := entity.SetI(10)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i)

}
