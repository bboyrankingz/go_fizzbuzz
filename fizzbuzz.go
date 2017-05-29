package main

import (
	"fmt"
	"strconv"
	"log"
	"net/http"
	"encoding/json"
)

type BuzzStrings []*string

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
	strings := BuzzStrings{}
	for i := 1; i < 100; i++ {
		fizz := WordOrEmpty(i, 3, "fizz")
		buzz := WordOrEmpty(i, 5, "buzz")
		//fmt.Println(i)
		fizzbuzz := fizz + string(buzz)
		if fizzbuzz != ""{
			strings = append(strings, &fizzbuzz)
		} else {
			t := strconv.Itoa(i)
			//fmt.Println(strings)
			strings = append(strings, &t)
		}
	}

	b, _ := json.Marshal(strings)

	fmt.Fprintf(writer, string(b))
}

func WordOrEmpty(i int, number int, word string) string {
	r := ""
	if i % number == 0 {
		r = word
	}
	return r
}
