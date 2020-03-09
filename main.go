package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "삼겹살"}
	superman := person{"suuper", 23, favFood}
	batman := person{name: "baat", age: 50, favFood: favFood}
	fmt.Println(superman, batman)
}
