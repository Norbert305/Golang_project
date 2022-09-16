package main

import "fmt"

type Pet struct {
	name    string
	petType string
	age     int
}

func main() {

	Mittens := Pet{"Mittens", "cat", 5}

	Buddy := Pet{"Buddy", "dog", 10}

	fmt.Println(Mittens)
	fmt.Println(Buddy)

}
