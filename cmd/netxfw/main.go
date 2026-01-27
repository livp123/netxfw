package main

import (
	"github.com/livp123/netxfw/internal/xdp"
	"log"
)

func run() {
	// ...
	interfaces, err := xdp.GetPhysicalInterfaces()
	if err != nil {
		log.Fatalf("Get interfaces: %v", err)
	}

	manager, err := xdp.NewManager()
	if err != nil {
		log.Fatalf("Create XDP manager: %v", err)
	}
	defer manager.Close()

	if err := manager.Attach(interfaces); err != nil {
		log.Fatalf("Attach XDP: %v", err)
	}

	// 封禁示例
	xdp.BanIP(manager.Blacklist(), "192.168.1.100")
}
