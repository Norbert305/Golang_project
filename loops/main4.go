letters := []string{"A", "B", "C", "D"}
for index, value := range letters {
  fmt.Println("Index:", index, "Value:", value)
}

//output 
Index: 0 Value: A
Index: 1 Value: B
Index: 2 Value: C
Index: 3 Value: D

addressBook := map[string]string{
	"John": "12 Main St",
	"Janet": "56 Pleasant St",
	"Jordan": "88 Liberty Ln",
  }
  for key, value := range addressBook {
	fmt.Println("Name:", key, "Address:", value)
  }
//output
Name: John Address: 12 Main St
Name: Janet Address: 56 Pleasant St
Name: Jordan Address: 88 Liberty Ln