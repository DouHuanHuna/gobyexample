package main

import "fmt"

func main() {

	map_ := map[string]int{
		"one": 1,
		"two": 2,
	}

	if _, ok := map_["three"]; !ok {
		fmt.Println("no three")
	}

}
