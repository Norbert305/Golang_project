package main

import "fmt"

//%T --> dataType
//%d --> Int
//%v --> string

func main() {
  floatExample := 1.75
  // Edit the following Printf for the FIRST step
  fmt.Printf("Working with a %T", floatExample) 
  
  fmt.Println("\n***") // Added for spacing
  
  myString := "hi"
  yearsOfExp := 3
  reqYearsExp := 15
  // Edit the following Printf for the SECOND step
  fmt.Printf("I have %d years of Go experience and this job is asking for %d years %v.", yearsOfExp, reqYearsExp, myString) 
  
  fmt.Println("\n***") // Added for spacing
  
  stockPrice := 3.50
  // Edit the following Printf for the THIRD step
  fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice) 
  fmt.Println("\n***")
  myInt := 5
  fmt.Printf("The data type being used is an %T", myInt)
}





