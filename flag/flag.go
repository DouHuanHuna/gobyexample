package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
)

func main() {
	metricsPath := kingpin.Flag(
		"web.telemetry-path",
		"Path under which to expose metrics.",
	).Default("/metrics").String()
	disableExporterMetrics := kingpin.Flag(
		"web.disable-exporter-metrics",
		"Exclude metrics about the exporter itself (promhttp_*, process_*, go_*).",
	).Bool()

	kingpin.Parse()
	fmt.Println(*metricsPath)
	if *disableExporterMetrics {
		fmt.Println("Exporter metrics are disabled.")
	} else {
		fmt.Println("Exporter metrics are enabled.")
	}
}
