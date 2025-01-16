package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	metricName := prometheus.BuildFQName("default", "wsl", "fq_name")
	fmt.Println(metricName)

}
