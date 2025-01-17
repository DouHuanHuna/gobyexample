package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {

	registry := prometheus.NewRegistry()

	requestsTotal := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "requests_total",           // TYPE {Name} {RegistryType}
		Help: "Total number of requests", // HELP {Name} Help
	})

	registry.MustRegister(requestsTotal)

	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	go func() {
		fmt.Println("Starting HTTP server on :8000")
		if err := http.ListenAndServe(":8000", nil); err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		requestsTotal.Inc()
		fmt.Println("Request count:", requestsTotal.Desc())
	}
}
