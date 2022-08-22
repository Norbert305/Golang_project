package main

import (
  "fmt"
)

func main() {

jellybeans := []string{"green", "blue", "yellow", "red", "green", "yellow", "red"}
for index := 0; index < len(jellybeans); index++ {
  if jellybeans[index] == "green" {
    continue //will continue to loop through the entire array excuding the green index
  }
  fmt.Println("You ate the", jellybeans[index], "jellybean!")
}


//------------------------------------------------------------------------------


animals := []string{"Cat", "Dog", "Fish", "Turtle"}
for index := 0; index < len(animals); index++ {
  if animals[index] == "Dog" {
    fmt.Println("Found the perfect animal!")
    break // will stop once the index is found based upon the condition
	//stops loop at dog
  }
}

}