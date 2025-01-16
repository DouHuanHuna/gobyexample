package main

import (
	"fmt"
	"github.com/prometheus/procfs"
	"os"
	"path/filepath"
)

func main() {
	data, err := os.ReadFile(filepath.Join(procfs.DefaultMountPoint, "loadavg"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
