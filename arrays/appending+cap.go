package main

import "fmt"

func main() {
    myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
    myTutors = append(myTutors, "Norbert")//output [Kirsty Mishell Jose Neil Norbert]
    fmt.Println(myTutors)
    fmt.Println(len(myTutors))//output 5
    fmt.Println(cap(myTutors))//output 8
}