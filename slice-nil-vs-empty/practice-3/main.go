package main

import "fmt"

func main() {
	s := order()
	fmt.Println(s)
}

func order() []string {
	var s []string
	if fruit() {
		s = append(s, "fruit")
	}
	if vegitable() {
		s = append(s, "vegitable")
	}
	return s
}

func fruit() bool {
	return true
}

func vegitable() bool {
	return true
}
