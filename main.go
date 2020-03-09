package main

import "fmt"

func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 19:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(20))
	fmt.Println(canIDrink(17))
}
