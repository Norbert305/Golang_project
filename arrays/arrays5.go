package main

import "fmt"

func main() {
  triangleAngles := [3]int{30, 60, 90}
  //Print largest angle number in array
  fmt.Println(triangleAngles[2])//output 90

  //sum number elelments inside of array using array indexing to store into sum variable
  sum := triangleAngles[0] + triangleAngles[1] + triangleAngles[2]//output 180

    //print sum variable
   fmt.Println(sum)
}