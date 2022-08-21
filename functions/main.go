package main

import "fmt"


//define functions below



func startGame() {
  instructions := "Press enter to start..." 
  
    fmt.Println(instructions)
  
}


func performAddition() {
  x := 5
  y := 7
  fmt.Println("The sum of", x, "and", y, "is", x + y)
}

func multiplyMe(a int, b int) int {
		return a * b
}

func main() {
	//call functions here!!!!!
  startGame()
    performAddition()
	fmt.Println(multiplyMe(5,5))
}