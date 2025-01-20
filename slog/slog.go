package main

import (
	"log/slog"
	"os"
)

func main() {
	handler := slog.NewTextHandler(os.Stdout, nil)

	slog.New(handler)

	slog.Info("slog hello")
	//
	//logger := slog.Logger{}
	//logger.Debug("slog hello")
}
