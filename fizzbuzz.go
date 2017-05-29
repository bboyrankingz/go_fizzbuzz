package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/fizzbuzz", fizzbuzzHandler)
}

func main() {
	address := "localhost:8080"
	log.Printf("Starting Server listening on %s\n", address)
	// create and start Server with DefaultServeMux Handler
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}

func fizzbuzzHandler(writer http.ResponseWriter, request *http.Request) {
	var a []string
	sum := 1
	for ; sum < 100; {
		fizz := WordOrEmpty(sum, 3, "fizz")
		buzz := WordOrEmpty(sum, 5, "buzz")
		sum += sum
		fizzbuzz := fizz + string(buzz)
		if fizzbuzz != ""{
			a = append(a, fizzbuzz)
		} else {
			a = append(a, string(sum))
		}
	}
	fmt.Fprintf(writer, "Hello Fizzbuzz")
}

func WordOrEmpty(i int, number int, word string) string {
	r := ""
	if i % number == 0 {
		r = word
	}
	return r
}
