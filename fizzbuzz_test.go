package main

import (
	"testing"
	"net/http/httptest"
	"io/ioutil"
	"encoding/json"
	"log"
)

func TestWordOrEmptyReturnHello(t *testing.T) {
	actualResult := WordOrEmpty(10, 5,"Hello")
	var expectedResult = "Hello"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestHelloWorldHttp(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)

	w := httptest.NewRecorder()
	helloHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Fatalf("Response status Expected %d but got %d", 200, resp.StatusCode)
	}
	var actualResult = string(body)
	var expectedResult = "Hello World"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestFizzBuzzdHttp(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080/fizzbuzz", nil)

	w := httptest.NewRecorder()
	fizzbuzzHandler(w, req)

	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Fatalf("Response status Expected %d but got %d", 200, resp.StatusCode)
	}
	actualResult := []string{}
	var expectedResult = []string{"1", "2", "fizz", "4", "buzz"}
	if err := json.NewDecoder(resp.Body).Decode(&actualResult); err != nil {
		log.Fatalln(err)
	}

	if len(actualResult) != 100 {
		t.Fatalf("Expected 100 elements but got %d", len(actualResult))
	}

	for index, element := range expectedResult {
		if element != actualResult[index] {
			t.Fatalf("Expected %s but got %s", actualResult[index], element)
		}
	}
}

func TestFizzBuzzWithParametersdHttp(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080/fizzbuzz?string1=foo&string2=bar&int1=2&int2=5&limit=50", nil)

	w := httptest.NewRecorder()
	fizzbuzzHandler(w, req)

	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Fatalf("Response status Expected %d but got %d", 200, resp.StatusCode)
	}
	actualResult := []string{}
	var expectedResult = []string{"1", "foo", "3", "foo", "bar"}
	if err := json.NewDecoder(resp.Body).Decode(&actualResult); err != nil {
		log.Fatalln(err)
	}

	if len(actualResult) != 50 {
		t.Fatalf("Expected 50 elements but got %d", len(actualResult))
	}

	for index, element := range expectedResult {
		if element != actualResult[index] {
			t.Fatalf("Expected %s but got %s", actualResult[index], element)
		}
	}
}