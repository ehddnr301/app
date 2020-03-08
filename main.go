package main

import "fmt"

func main() {
	const name string = "hello"

	var studying bool = true
	studying = false

	shortcut := "good"
	shortcut = "excellent"

	fmt.Println(name, studying, shortcut)
}
