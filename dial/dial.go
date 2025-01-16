package main

import (
	"fmt"
	"github.com/mdlayher/netlink"
	"syscall"
)

func main() {
	conn, err := netlink.Dial(syscall.NETLINK_INET_DIAG, nil)
	if err != nil {
		return
	}

	fmt.Println(conn)
}
