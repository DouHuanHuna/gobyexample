package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "你好，世界！世界很美好！"
	s = strings.ReplaceAll(s, "世界", "地球")
	fmt.Println(s)
}
