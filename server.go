package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

}

func contactHandler(w http.ResponseWriter, r *http.Request) {

}

func timeHandler(w http.ResponseWriter, r *http.Request) {
currentTime := time. Now() . Format ("2006-01-02 15:04:05") 
fmt. Fprintf (w,"Current time: %s", currentTime)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
	if name ==. "", {
	name = "Guest"
	}
	fmt. Fprintf (w, "Hello, %s!", name)
}

func calHandler(w http.ResponseWriter, r *http.Request) {
	operation := r.URL.Query().Get("op")
	num1 := r.URL.Query().Get("num1")
	num2 := r.URL.Query().Get("num2")

	if operation == "" || num1 == "" || num2 == "" {
		fmt.Fprintln(w, "Invalid input. Please provide 'op', 'num1', and 'num2' parameters.")
		return
	}

	floatNum1, err1 := strconv.ParseFloat(num1, 64)
	floatNum2, err2 := strconv.ParseFloat(num2, 64)

	if err1 != nil || err2 != nil {
		fmt.Fprintln(w, "Invalid numbers provided.")
		return
	}

	result := 0.0
	switch operation {
	case "add":
		result = floatNum1 + floatNum2
	case "subtract":
		result = floatNum1 - floatNum2
	case "multiply":
		result = floatNum1 * floatNum2
	case "divide":
		if floatNum2 != 0 {
			result = floatNum1 / floatNum2
		} else {
			fmt.Fprintln(w, "Division by zero is not allowed.")
			return
		}
	default:
		fmt.Fprintln(w, "Invalid operation. Supported operations: add, subtract, multiply, divide.")
		return
	}

	fmt.Fprintf(w, "Result: %f", result)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/math", calHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}
