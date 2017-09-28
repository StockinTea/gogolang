package main

import "fmt"

func receiveMan(Msg chan<- string, PassMsg string) {
  Msg <- PassMsg
}

func sendMan(PassMsg <-chan string, SendMsg chan<- string) {
  msg := <- PassMsg
  SendMsg <- msg
}

func badMan(SendMsg <-chan string) {
  fakeMsg := "Jeff has new IPad"
  SendMsg <- fakeMsg
}

func main() {
  msg := make(chan string, 1)
  finialMsg := make(chan string, 1)

  go receiveMan(msg, "Devin has new IPad")
  go sendMan(msg, finialMsg)
  go badMan(finialMsg)

  fmt.Println(<-finialMsg)
}
