package main

import "fmt"

func main() {
	superman := map[string]string{"name": "super", "age": "12"}
	for key, value := range superman {
		fmt.Println(key, value)
	}
}
