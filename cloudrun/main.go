package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		numParam := r.URL.Query().Get("num")
		fmt.Println("Received request: ", numParam)

		num, err := strconv.ParseInt(numParam, 10, 64)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		fmt.Fprintf(w, fmt.Sprint(fibonacci(num)))
	})

	http.ListenAndServe(":8080", nil)
}
