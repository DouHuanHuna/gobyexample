package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

type handler struct {
	httpHandler http.Handler
	logger      *slog.Logger
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	collects := r.URL.Query()["collect[]"]
	fmt.Printf("collects: %v\n", collects)
}
func main() {

	h := &handler{}

	http.HandleFunc("/", h.ServeHTTP)
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		return
	}
}
