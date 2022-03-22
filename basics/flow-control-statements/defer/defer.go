package main

import "fmt"

func main() {
	defer fmt.Print("world.\n")
	fmt.Print("Hello ")
}
