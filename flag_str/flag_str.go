package main

import "flag"

func main() {
	addr := flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()
	println("addr:", *addr)
}
