package main

import (
	"fmt"
	"github.com/prometheus/procfs"
	"log"
)

func main() {
	fs, err := procfs.NewFS("/proc")
	if err != nil {
		log.Fatalf("failed to open procfs: %v", err)
	}

	sockStat, err := fs.NetSockstat()
	if err != nil {
		log.Fatalf("failed to get IPv4 sockstat data: %v", err)
	}

	fmt.Printf("IPv4 Sockets Used: %d\n", *sockStat.Used)
	for _, protocol := range sockStat.Protocols {
		fmt.Printf("Protocol: %s\n", protocol.Protocol)
		fmt.Printf("InUse: %d\n", protocol.InUse)
		if protocol.Orphan != nil {
			fmt.Printf("Orphan: %d\n", *protocol.Orphan)
		}
		if protocol.TW != nil {
			fmt.Printf("TW: %d\n", *protocol.TW)
		}
		if protocol.Alloc != nil {
			fmt.Printf("Alloc: %d\n", *protocol.Alloc)
		}
		if protocol.Mem != nil {
			fmt.Printf("Mem: %d\n", *protocol.Mem)
		}
		if protocol.Memory != nil {
			fmt.Printf("Memory: %d\n", *protocol.Memory)
		}
	}
}
