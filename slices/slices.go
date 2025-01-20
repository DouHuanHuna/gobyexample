package main

import "fmt"

func main() {
	stats := []struct {
		name string
		age  int
	}{
		{"John Smith", 25},
		{"John Smith", 26},
		{"John Smith", 27},
	}

	fmt.Println(stats)
	stats = append(stats, struct {
		name string
		age  int
	}{
		"git", 11,
	})

	fmt.Println(stats)
}
