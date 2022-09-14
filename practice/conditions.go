package main

import "fmt"

func teacher() {

	pass := "hello"
	pass2 := "hello"

	if pass == pass2 {
		fmt.Println("suceess 200")
	} else {
		fmt.Println("Bad 400")
	}

}

func main() {

	teacher()

}
