package main
import "fmt"

func main() {
  donuts := map[string]int{
    "frosted":   10,
    "chocolate": 15,
    "jelly":     8,
  }
  fmt.Println(donuts)//map[chocolate:15 frosted:10 jelly:8]

  // Add your code here
  delete(donuts, "chocolate")

  fmt.Println(donuts)//map[frosted:10 jelly:8]
}