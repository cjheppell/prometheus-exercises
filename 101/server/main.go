package main

import (
	"fmt"
	"net/http"
)

func counter(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from the counter route\n")
}

func gauge(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from the gauge route\n")
}

func histogram(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from the histogram route\n")
}

func main() {
	http.HandleFunc("/counter", counter)
	http.HandleFunc("/gauge", gauge)
	http.HandleFunc("/histogram", histogram)

	http.ListenAndServe(":8080", nil)
}
