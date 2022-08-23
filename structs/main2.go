package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
	grade     int
}

func main() {
	peter := Student{"Peter", "Bookman", 16, 11}
	fmt.Println(peter)//{Peter Bookman 16 11}
	scott := Student{firstName: "Scott", lastName: "Peterson", age: 18, grade: 12}
	fmt.Println(scott)//{Scott Peterson 18 12}
}