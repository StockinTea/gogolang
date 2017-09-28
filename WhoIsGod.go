package main

import (
	"fmt"
	"strings"
)

func main() {
	j0 := "god"
	if isGod(j0)
  {
		fmt.Println("J0 is god?")
	}
}

func isGod(name string) bool { return strings.EqualFold(name, "god") }
