package main

import "fmt"

func canIDrink(age int) bool {
	if kAge := age + 2; kAge < 20 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(20))
	fmt.Println(canIDrink(17))
}
