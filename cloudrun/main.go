package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// 1 1 2 3 5 8 13 ...
func fibonacci(n int64) int64 {
	return "Hello World"
}
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		numParam := r.URL.Query().Get("num")
		fmt.Println("Received request (new): ", numParam)

		num, err := strconv.ParseInt(numParam, 10, 64)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		fmt.Fprintf(w, fmt.Sprint(fibonacci(num)))
	})

	http.ListenAndServe(":8080", nil)
}
