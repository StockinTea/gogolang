package main

import (
	"fmt"
	"sync"
)
var wait sync.WaitGroup

func main() {
  wait.Add(1)
	go threadAndyYell("Hey!")

	fmt.Println("Go Go Andy")

	wait.Wait()
	fmt.Println("done")
}

func threadAndyYell(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str)
	}
	wait.Done()
}
