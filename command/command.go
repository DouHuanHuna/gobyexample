package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "-c", "echo hello world")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))

}
