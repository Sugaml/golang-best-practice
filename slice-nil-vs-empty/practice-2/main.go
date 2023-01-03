package main

import "fmt"

func main() {
	s1 := make([]int, 0)
	s1 = append(s1, 1) // okay

	fmt.Println(s1)

	var s2 []int
	s2 = append(s2, 1) //[1]
	fmt.Println(s2)

}
