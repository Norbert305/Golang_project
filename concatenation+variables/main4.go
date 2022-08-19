package main //declare our package

import "fmt" //using fmt package to print out our data

func main () {

var publisher string = "DizzyBooks Publishing Inc"
var writer string = "Tracy Hatchet"
var artist string = "Jewel Tampson"
var title string = "Mr. GoToSleep"

var year int = 1997
var pageNumber int = 14

var grade float32 = 6.5

fmt.Println(title, "written by", writer, "drawn by", artist, "year came out was", year, "number of pages was", pageNumber, "and the final score was a", grade, "rating and was published by", publisher)

// Assigns will go here
 title = "Epic Vol. 1"
writer = "Ryan N. Shawn"
artist = "Phoebe Paperclips"
publisher = "DizzyBooks Publishing Inc."
year = 2013
pageNumber = 160
grade = 9.0


fmt.Println(title, "written by", writer, "drawn by", artist, "year came out was", year, "number of pages was", pageNumber, "and the final score was a", grade, "rating and was published by", publisher)

}

