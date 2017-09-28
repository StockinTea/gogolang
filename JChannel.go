package main

import (
	"fmt"
	"time"
)

func j_text(str chan string, msg string) {
	for i := 0; ; i++ {
		str <- msg
	}
}
func j_print(str chan string) {
	for {
		fmt.Println(<-str)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	const ti = 5
  str := make(chan string)

	go j_text(str, "Do you like j_channel ?")
	go j_print(str)

	time.Sleep(time.Second * ti)
	fmt.Println("done")
}
