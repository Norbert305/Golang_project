package main
import (
  "fmt"
)

func add(a int, b int) int {
  return a + b
}

func printMessage (age int, name string, dessert string) {
    fmt.Println("My name is",name,"and I am",age,"years old and I love", dessert,"!")
}

func main() {
    fmt.Println(add(4,5))
    printMessage(31,"Norbert", "icecream")
}