package main

import (
  "fmt"
)

func main() {
var name string
fmt.Println("Hello, what's your name?")
fmt.Scan(&name)
fmt.Println("Welcome",name)
var age int
fmt.Println("What's your age?")
fmt.Scan(&age)
fmt.Println("I am",age,"years old!")
fmt.Printf("My name is %s and I am %d years old!!!",name,age)

newMessage := fmt.Sprintf("My actual name is %s and my real age is %d", name, age)

fmt.Println(newMessage)

}