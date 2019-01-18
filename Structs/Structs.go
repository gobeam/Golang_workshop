package main

import "fmt"

type company struct {
	name      string
	address   string
	employees int
}

type person struct {
	name   string
	office company
}

func main() {
	company := company{
		name:      "Ekbana",
		address:   "Jwagal",
		employees: 100,
	}
	fmt.Println(company)

	person := person{
		name: "John Doe",
		office: company,
	}
	fmt.Println(person)
}
