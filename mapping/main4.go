package main
import "fmt"

func main() {
  donuts := map[string]int{
    "frosted":   10,
    "chocolate": 15,
    "jelly":     8,
  }

  // Print out all the donuts
  fmt.Println(donuts) //map[chocolate:15 frosted:10 jelly:8]

  // Add your code here
  donuts["glazed"] = 12

  fmt.Println(donuts["glazed"])//12

  donuts["jelly"] = 3
  fmt.Println(donuts)//map[chocolate:15 frosted:10 glazed:12 jelly:3]
}