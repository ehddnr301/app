package main

import "fmt"

func main() {
	names := [5]string{"as", "ax", "ew"}
	supers := []string{"as123", "ax123", "ew123"}
	supers = append(supers, "superman")
	fmt.Println(names, supers)
}
