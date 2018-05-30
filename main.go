package main

import "fmt"

//Customer bla bla
type Customer struct {
	firstName string
	lastName  string
}

func main() {
	customer := Customer{"Savas", "Ozturk"}
	_ = customer

	customer2 := Customer{}
	customer2.firstName = "Derin"
	customer2.lastName = "Ozturk"

	n, err := fmt.Printf("customerName:%v\n", customer2)
	if err == nil {
		fmt.Printf("n:%+v", n)
	}

	helloWorld("Savas Ozturk")

}
