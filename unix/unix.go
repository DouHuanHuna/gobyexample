package main

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func main() {
	path := "/"
	var stat unix.Statfs_t

	err := unix.Statfs(path, &stat)
	if err != nil {
		fmt.Println("Something went wrong", err)
	}
	fmt.Println(stat)

}
