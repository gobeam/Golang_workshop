package main

import (
	"Golang_workshop/Package/Add"
	"fmt"
)

func main() {
//package declaration must be on top of page

result := Add.Add(4,5)

fmt.Println("/n addition result is: ", result)

}
