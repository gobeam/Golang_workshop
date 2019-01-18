package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// without using Goroutine
	start := time.Now()
	for a := 0; a < 5; a++ {
		callApi("http://pokeapi.co/api/v2/pokedex/kanto/")
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Elapsed time without using go routine is: ",elapsed)


	start2 := time.Now()
	resp := make(chan string, 5)
	for a := 0; a < 5; a++ {
		go callApiWithRoutine("http://pokeapi.co/api/v2/pokedex/kanto/", resp)
	}
	for a := 0; a < 5; a++ {
		<-resp
	}
	t2 := time.Now()
	elapsed2 := t2.Sub(start2)

	fmt.Println("Elapsed time with using channels is: ",elapsed2)
}

func callApi(url string) string {
	//response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(responseData)
}

func callApiWithRoutine(url string,resp chan<- string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp <- string(responseData)
}
