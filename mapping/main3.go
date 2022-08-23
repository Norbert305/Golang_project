package main
import "fmt"

func main() {
  donuts := map[string]int{
    "frosted":   10,
    "chocolate": 15,
    "jelly":     8,
  }
   
  // Print out the inventory
  fmt.Println(donuts) //map[chocolate:15 frosted:10 jelly:8]
   
  // Add your code here
  firstChoice := donuts["frosted"]
  fmt.Println(firstChoice)//10

  secondChoice, status := donuts["bavarian cream"]

  fmt.Println(secondChoice)//0 
  fmt.Println(status)  //false

  //if statement will print false statement
  if status {
    fmt.Println(secondChoice)
  } else {
    //will print else statement
    fmt.Println("We do not sell that donut!")
  }

}