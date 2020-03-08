package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, upper string) {
	defer fmt.Println("It's Done")
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

func main() {
	totalLen, uppername := lenAndUpper("Someting")
	fmt.Println(totalLen, uppername)
}
