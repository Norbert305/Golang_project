package main

import "fmt"

func main() {
  // Add your code here
  orders := make(map[string]float32)
  fmt.Println(orders) //map[]
  //output empty map
 

  donuts := map[string]int{
    "frosted" : 10,
    "chocolate" : 15,
    "jelly" : 8,
  }
  fmt.Println(donuts) //map[chocolate:15 frosted:10 jelly:8]
}