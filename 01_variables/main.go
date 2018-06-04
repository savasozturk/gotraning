package main

import "fmt"

//Customer comment not to throw warning!
type Customer struct {
	firstName string
	lastName  string
}

func main() {
	customer := Customer{"Savas", "Ozturk"}
	n, err := fmt.Printf("customerName:%v\n", customer)
	if err == nil {
		fmt.Printf("n:%+v", n)
	}
}
