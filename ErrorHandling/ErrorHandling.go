package main

import (
	"errors"
	"fmt"
)

func main(){

	// way of handling error
	val, err := errorExample(3)
	if err != nil {
		msg := fmt.Sprintf("%s", err.Error())
		fmt.Println("\n ", msg)
	}

	fmt.Println("\n ", val)

}


func errorExample(val int)(string, error) {
	if val > 2 {
		err := errors.New("value is greater than 2, please provide smaller value")
		return "Bad boi no treat from hooman", err
	}else {
		return "Good boi, hooman will get treat for you!", nil
	}
}
