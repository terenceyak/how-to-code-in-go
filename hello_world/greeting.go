package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please enter your name.")
	var name string
	name = strings.TrimSpace(name)
	fmt.Scanln(&name)
	fmt.Printf("Hi, %s! I'm GO!", name)
}
