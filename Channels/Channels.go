package main

import (
	"fmt"
	"time"
)

func main()  {
	// Channels provide us concurrent mechanism to execute function which can send or receive value of certain type
	// Channel has <- operator which specifies if channels can receive or send value
	// If there is no operator <- in channel then channel can receive or send which means it is bidirectional

	//channel can be initialize with make
	// Below channel can receive or send string value only unbuffered channel because we haven't specified size this channel can hold an
	newChan := make(chan string)


	go message(newChan)
	msg := <-newChan
	fmt.Println(msg)

	// bad way to use channels
	go badWayToUseChannel()

	// workaround which is still weird
	time.Sleep(time.Second)
	
}


func badWayToUseChannel() {
	fmt.Println("Oh my god when will my message print......")
}

func message(newChan chan string) {
	newChan <- "you got message my man!"
}
