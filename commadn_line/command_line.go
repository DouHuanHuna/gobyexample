package main

import (
	"fmt"
	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/exporter-toolkit/web/kingpinflag"
)

func main() {
	toolkitFlags := kingpinflag.AddFlags(kingpin.CommandLine, ":9100")
	kingpin.Parse()
	fmt.Println(*toolkitFlags)
}
