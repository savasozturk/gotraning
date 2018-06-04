package main

import "fmt"

//Customer comment not to throw warning!
type Customer struct {
	id        int
	firstName string
	lastName  string
}

func main() {
	customer := Customer{7, "Savas", "Ozturk"}
	n, err := fmt.Printf("customerName:%v\n", customer)
	if err == nil {
		fmt.Printf("n:%+v", n)
	}
}
