package main

import "fmt"

func main() {

	addressBook := map[string]string{
		"John":   "12 Main St",
		"Janet":  "56 S Bay Ave",
		"Jordan": "23 Bulls Drive",
	}

	for key, value := range addressBook {
		fmt.Println("Name:", key, "Address:", value)
	}

}
