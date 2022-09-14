package main

import "fmt"

func main() {

	animals := []string{"cat", "dog", "mouse", "tiger"}

	for i := 0; i < len(animals); i++ {
		if animals[i] == "dog" {
			fmt.Println("I found the dog")
			break
		}
	}

}
