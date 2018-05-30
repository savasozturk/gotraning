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
	fmt.Printf("Customer type print:%v\n", customer)
	fmt.Printf("Customer name:%v", customer.firstName)
}
