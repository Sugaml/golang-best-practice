package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Name        string
	ContactInfo []string
}

func main() {
	//nil slice
	var s1 []string
	customer1 := employee{
		Name:        "employee-1",
		ContactInfo: s1,
	}
	b, _ := json.Marshal(customer1)
	fmt.Println(string(b))

	//non-nil empty slice
	s2 := make([]string, 0)
	customer2 := employee{
		Name:        "employee-2",
		ContactInfo: s2,
	}
	b, _ = json.Marshal(customer2)
	fmt.Println(string(b))
}
