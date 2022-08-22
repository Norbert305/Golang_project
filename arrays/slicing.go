package main

import "fmt"

func main() {
    myTutors := [4]string{"Kirsty", "Mishell", "Jose", "Neil"}
    
      fmt.Println(myTutors)
            //theory 1 = stop index | 3 = stop length
    sliceVersion := myTutors[1:3] //output [Mishell Jose]

    subjects := []string{"Go", "Java", "Python"}
              //theory 2 leftside only = stop at index 
    sliceVersion2 := subjects[2:] //output [Pyhton]
    // sliceVersion2 := subjects[:1] //output [Go]
    // sliceVersion2 := subjects[:2] //output [Go Java]

    fmt.Println(sliceVersion)
    fmt.Println(sliceVersion2)
    fmt.Println(subjects)
}