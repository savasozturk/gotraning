package main

import "fmt"

//Customer bla bla
type Customer struct {
	firstName string
	lastName  string
}

func main() {

	customer := Customer{"Savas", "Ozturk"}
	_ = customer //to override warning!

	customer2 := Customer{}
	customer2.firstName = "Derin"
	customer2.lastName = "Ozturk"

	n, err := fmt.Printf("customerName:%v\n", customer2)
	if err == nil {
		fmt.Printf("n:%+v", n)
	}
}
