package main

import "fmt"

type Restaurant struct {
	name             string
	typeOfRestaurant string
	yearEstablished  int
}

func main() {

	// Add your code here.
  restaurant := Restaurant{"Codecademy Steakhouse", "Japanese", 2011}
  fmt.Println(restaurant)//{Codecademy Steakhouse Japanese 2011}

  fmt.Println(restaurant.name)//Codecademy Steakhouse
  fmt.Println(restaurant.typeOfRestaurant)//Japanese
  fmt.Println(restaurant.yearEstablished)//2011

  restaurant.name = "Skillsoft Steakhouse"
  restaurant.yearEstablished = 2022
  fmt.Println(restaurant)//{Skillsoft Steakhouse Japanese 2022}
}