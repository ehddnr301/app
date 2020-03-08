package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func consoleLogging(name ...string) {
	fmt.Println(name)
}

func main() {
	totalLen, uppername := lenAndUpper("Someting")
	totalLength, _ := lenAndUpper("sssssssss")
	fmt.Println(totalLen, uppername, totalLength)

	consoleLogging("super", "suppppper", "many", "arguments")
}
