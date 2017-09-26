package main

import (
	"fmt"
	"strings"
)

func main() {
	j0 := "god"
	if isGod(j0) {
		fmt.Println("Hello, 露天")
	}
}

func isGod(name string) bool { return strings.EqualFold(name, "god") }
