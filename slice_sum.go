package main

func main()  {
	items := []int{10, 20, 30, 40, 50}
	fmt.Println(items)
	sum := sum(items...)
	fmt.Printf("The sum of items is equal to %d.\n", sum)	
}
func sum(items ...int) int {
	sum := 0
	for _, item := range items {
		sum += item
	}
	return sum
}
