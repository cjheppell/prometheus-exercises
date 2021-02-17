package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counterRequestsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "prom101_server_counter_requests_total",
		Help: "The total number of processed counter requests events",
	})
)

func counter(w http.ResponseWriter, req *http.Request) {
	counterRequestsProcessed.Inc()
	fmt.Fprintf(w, "Hello from the counter route\n")
}

func guageStart(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from the gauge start route\n")
}

func guageStop(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from the gauge stop route\n")
}

func histogram(w http.ResponseWriter, req *http.Request) {
	delayTime := rand.Intn(10)
	time.Sleep(time.Duration(delayTime) * time.Second)
	fmt.Fprintf(w, "Hello from the histogram route\n")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/counter", counter)
	http.HandleFunc("/start", guageStart)
	http.HandleFunc("/stop", guageStop)
	http.HandleFunc("/histogram", histogram)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
