package tun

import (
	"log"

	"github.com/songgao/water"
)

func CreateTUN() (*water.Interface, error) {
	config := water.Config{
		DeviceType: water.TUN,
	}
	config.Name = "tun0"

	ifce, err := water.New(config)
	if err != nil {
		return nil, err
	}

	log.Printf("TUN interface created: %s\n", ifce.Name())
	return ifce, nil
}
