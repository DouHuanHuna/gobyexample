package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.Int("port", 8080, "")
	//------------------------------------------------//
	var name string
	flag.StringVar(&name, "server-name", "nginx", "")

	flag.Parse()
	fmt.Println("port:", *port, "name:", name)

}
