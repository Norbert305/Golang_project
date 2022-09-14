package main

import "fmt"

func main() {

	dog := "dog"
	cat := "cat"
	var bird string = "bird"

	compliment := "Good Job"

	teacher := fmt.Sprint("The teacher says", compliment, "on todays quiz!!!")

	fmt.Println("I have a", dog, "and a", cat, "in my house")
	fmt.Printf("Plus I have a %v", bird)
	fmt.Print(teacher)
}
