package main

import (
	"fmt"
	"sync"
)

func main() {
	var goSync sync.WaitGroup
	for i := 0; i < 10; i ++ {
		i1 := i
		goSync.Add(1)
		go calPrint(&i1, &goSync)
	}
	goSync.Wait()
}

func calPrint(target *int, goSync *sync.WaitGroup){
	fmt.Println(*target)
	defer goSync.Done()
}
