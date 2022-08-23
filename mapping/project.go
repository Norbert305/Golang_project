package main

import (
	"fmt"
)

func main() {
  var currencies = map[string]float32{"JPY" : 130.2, "EUR" : 0.95,}

  var dollarAmount float32
  var currency string
  fmt.Println("Please enter dollar amount")
  fmt.Scan(&dollarAmount)
  if dollarAmount == 0 {
  fmt.Println("Not enough here!!!!") 
  } else {
      fmt.Println("Thankyou come again :)")
    }

    fmt.Println("What is the target currency?")
    fmt.Scan(&currency)

    rate, isValid := currencies[currency]

    if isValid {
      fmt.Println(dollarAmount, "USD currency converts to", rate*dollarAmount, currency)
    } else {
      fmt.Println("Cannot convert to", currency, "currency")
    }
}