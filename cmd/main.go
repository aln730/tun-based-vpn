package main

import (
	"log"

	"github.com/aln730/tun-based-vpn/internal/tun"
)

func main() {
	ifce, err := tun.CreateTUN()
	if err != nil {
		log.Fatalf("Error creating TUN interface: %v", err)
	}
	defer ifce.Close()

	buf := make([]byte, 1500) // typical MTU
	for {
		n, err := ifce.Read(buf)
		if err != nil {
			log.Fatalf("Read error: %v", err)
		}
		log.Printf("Read %d bytes from TUN: %x\n", n, buf[:n])
	}
}
