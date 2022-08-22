package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Println("You have", fuel, "gallons of fuel left!")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
    var fuel int
    if planet == "Mercury" {
  fuel = 500000
} else if planet == "Venus" {
  fuel = 300000
} else if planet == "Mars" {
  fuel = 700000
}
return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {

fmt.Println("Welcome to planet", planet)

}

// Create the function cantFly() here
func canFly() {

fmt.Println("We do not have the available fuel to fly there.")

}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
         var fuelRemaining int
         var fuelCost int
         fuelRemaining = fuel 
      fuelCost = calculateFuel(planet)

      if fuelRemaining >= fuelCost {
  greetPlanet(planet)
  fuelRemaining -= fuelCost
} else {
  canFly()
}
return fuelRemaining
}

func main() {
  // Test your functions!
  fmt.Println(calculateFuel("Venus"))
  greetPlanet("Venus")
  fuelGauge(5)

  // Create `planetChoice` and `fuel`
  var fuel int = 1000000
  var planetChoice string = "Venus"
  fuel = flyToPlanet(planetChoice, fuel)
    fuelGauge(fuel)
  // And then liftoff!
  
}