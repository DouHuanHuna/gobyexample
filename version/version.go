package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
	"github.com/prometheus/common/version"
)

func main() {
	application := kingpin.Version(version.Print("node_exporter"))
	kingpin.Parse()
	fmt.Println(*application)
}
