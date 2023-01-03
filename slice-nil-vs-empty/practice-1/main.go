package main

import "fmt"

func main() {
	//case 1
	var s []string
	log(1, s)

	//case 2
	s = []string(nil)
	log(2, s)

	//case3
	s = []string{}
	log(3, s)

	//case4
	s = make([]string, 0)
	log(4, s)
}

func log(i int, s []string) {
	fmt.Println("-------------------------------------------------")
	fmt.Printf("slice case %d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)
}
