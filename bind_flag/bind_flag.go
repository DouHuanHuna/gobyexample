package main

import (
	"flag"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func main() {
	var isEnabled bool

	flag.BoolVar(&isEnabled, "enable", false, "enable flag")

	opts := zap.Options{
		Development: false,
	}
	opts.BindFlags(flag.CommandLine)

	flag.Parse()

}
