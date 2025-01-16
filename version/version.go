package main

import (
	"github.com/alecthomas/kingpin"
	"github.com/prometheus/common/version"
)

func main() {
	kingpin.Version(version.Print("node_exporter"))
	kingpin.Parse()
}
