array := [5]int{2, 5, 7, 1, 3}
// This is a slice of the whole array
sliceVersion := array[:]
fmt.Println(sliceVersion)
// [2 5 7 1 3]
// This is a slice containing the elements at indices 2, 3, and 4
partialSlice := array[2:5]
fmt.Println(partialSlice)
// [7 1 3]


//-------------------------------------------------------------------------Access and change name elements




var names = []string{"Kathryn", "Martin", "Sasha", "Steven"}
fmt.Println(names[1])
// Martin
names[3] = "Bishop"
fmt.Println(names[3])
// Bishop