package main

import "fmt"

type company struct {
	name      string
	address   string
	employees int
}

type person struct {
	name string
	company
}

func main() {

	// Struct is collection field
	// Struct can be grouped together to form data

	//Empty struct example
	type Empty struct{}
	var e Empty
	fmt.Println("empty struct: ", e)

	// Struct with field
	company := company{
		name:      "Ekbana",
		address:   "Jwagal",
		employees: 100,
	}
	fmt.Println("\n company: ", company)

	// structs can be nested
	person := person{
		name:    "John Doe",
		company: company,
	}
	fmt.Println("\n person detail: ", person)

	// Elements can be accessed by:
	fmt.Println("\n Person name is: ", person.name)
	fmt.Println("\n Person works in: ", person.company.name, " located in ", person.company.address)

	//We can define our own function for structs
	person.printDetail()
}

// extra parameter appears before the function name is called method receiver
// In golang we dont have this or self to access object
func (p person) printDetail() {
	fmt.Println("\n Your name is ", p.name, ". Your office is located in ", p.company.address, " which contains total employee of ", p.company.employees)
}
