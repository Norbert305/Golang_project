package main

import "fmt"

func main() {

	lottery := [5]int{20, 50, 60, 30, 100}

	for i := 0; i < len(lottery); i++ {
		if lottery[i] == 50 {
			fmt.Println("I found number 50")
			break
		}
	}

	fmt.Println(lottery)

}
