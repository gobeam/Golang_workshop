package main

import "fmt"

type bike struct {
	ModelName string
	Mileage int
	Weight int
}

type car struct {
	ModelName string
	Mileage int
	Weight int
	HorsePower int
}

type vehicle interface {
	mileage()(string)
}

func main() {
	//same like any other programming language interfaces are named collection of set of method signatures
	// To implement interface we need to implement all the method defined in interface


	// Why interface?
	// It makes code more flexible and scalable since golang doesnot have concept of oops, it helps to achieve polymorphism

	//Example, lets take example of two different struct namely bike and car
	bike := bike{
		ModelName: "Aprilla",
		Mileage: 125,
		Weight: 100,
	}


	car := car{
		ModelName: "Mercedes",
		Mileage: 300,
		Weight: 200,
		HorsePower: 150,
	}

	//here both bike and car share similar method signature mileage as interface vehicle contain now rather than calling two separate method to print its details we defince on single method with vehicle interface as method input and it will automatically calls its corresponding method.
	printDetail(bike)
	printDetail(car)
}

func (b bike) mileage()(string) {
	return fmt.Sprintf("This bike has %d km/hr mileage",b.Mileage)
}

func (c car) mileage()(string) {
	return fmt.Sprintf("This car has %d km/hr mileage",c.Mileage)
}


func printDetail(v vehicle) {
	fmt.Println("\n Vehicle detail: ", v.mileage())
}
