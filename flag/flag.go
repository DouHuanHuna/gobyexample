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
	kingpin.Parse()
	fmt.Println(*metricsPath)
}
