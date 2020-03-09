package main

import "fmt"

func main1() {
	a := 2
	b := a
	c := &a
	a = 3
	fmt.Println(a, b, *c)
}

func main() {
	a := 2
	c := &a
	*c = 3
	fmt.Println(a)
}
