package main
import "fmt"



func changeLastElement(people []string, newName string){
    people[len(people)-1] = newName//last index of the array stored into variable called newName.//newName being changed to Keith
    fmt.Println(people)//arrayName of choice printed out
}

func main() {
    myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
    //My array of names
    changeLastElement(myTutors, "Keith")
    //function 1st arguemnt stores array.
    //function 2nd argument stores name of choice
}