package helloworld

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func init() {
	functions.HTTP("fibonacci", fibonacciHandler)
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	numParam := r.URL.Query().Get("num")
	num, err := strconv.ParseInt(numParam, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	fmt.Fprintf(w, fmt.Sprint(fibonacci(num)))
}
