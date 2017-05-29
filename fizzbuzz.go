package main

import (
	"fmt"
	"strconv"
	"log"
	"net/http"
	"net/url"
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
	u, err := url.Parse(request.URL.String())
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	number1 := getQueryInt(q.Get("int1"), 3)
	number2 := getQueryInt(q.Get("int2"), 5)
	limit := getQueryInt(q.Get("limit"), 100)
	string1 := getQueryString(q.Get("string1"), "fizz")
	string2 := getQueryString(q.Get("string2"), "buzz")

	strings := BuzzStrings{}
	for i := 1; i < limit + 1; i++ {
		fizz := WordOrEmpty(i, number1, string1)
		buzz := WordOrEmpty(i, number2, string2)
		fizzbuzz := fizz + string(buzz)
		if fizzbuzz != "" {
			strings = append(strings, &fizzbuzz)
		} else {
			t := strconv.Itoa(i)
			strings = append(strings, &t)
		}
	}

	b, _ := json.Marshal(strings)

	fmt.Fprintf(writer, string(b))
}
func getQueryInt(queryInt string, defaultValue int) int {
	number1 := defaultValue
	if queryInt != "" {
		number1, _ = strconv.Atoi(queryInt)
	}
	return number1
}

func getQueryString(queryInt string, defaultValue string) string {
	str1 := defaultValue
	if queryInt != "" {
		str1 = queryInt
	}
	return str1
}

func WordOrEmpty(i int, number int, word string) string {
	r := ""
	if i % number == 0 {
		r = word
	}
	return r
}
