package main

import "fmt"

func main() {
	slice := make([]int, 3)
	fmt.Println(slice)

	boolMake := make(map[string]*bool)
	fmt.Println(boolMake)

	b := true
	boolMake["true"] = &b
	fmt.Println(*boolMake["true"])
}
