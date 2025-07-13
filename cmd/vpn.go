package main

import (
	"crypto/rand"
	"log"

	"github.com/songgao/water"

	"github.com/aln730/tun-based-vpn/internal/crypto"
	"github.com/aln730/tun-based-vpn/internal/netpipe"
	"github.com/aln730/tun-based-vpn/internal/router"
)

func main() {
	// 1. Create TUN interface
	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatalf("TUN error: %v", err)
	}
	log.Printf("TUN interface: %s", ifce.Name())

	// 2. Generate key and create box
	key := make([]byte, 32)
	rand.Read(key)
	box, err := crypto.NewBox(key)
	if err != nil {
		log.Fatalf("crypto: %v", err)
	}

	// 3. Setup Pipe
	pipe, err := netpipe.NewPipe("127.0.0.1:9000", "127.0.0.1:9001", box)
	if err != nil {
		log.Fatalf("pipe: %v", err)
	}

	// 4. Start router
	r := router.New(ifce, pipe)
	r.Start()

	select {}
}
