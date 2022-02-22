package main

import "fmt"

func main() {
	arrray1 := [3]int{1, 2, 3}
	var slice1 []int
	//arrray1 = append(arrray1,4,5,6)
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(arrray1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)
}
