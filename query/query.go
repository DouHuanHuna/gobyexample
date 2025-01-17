package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	collectors := r.URL.Query()["collect[]"]
	if len(collectors) != 0 {
		fmt.Println("w:", w, "collectors:", collectors)
	}
}

func main() {
	//http://localhost:8089/?collect[]=value1&collect[]=value2&collect[]=value3
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8089", nil)

}
