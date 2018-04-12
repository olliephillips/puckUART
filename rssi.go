package puckUART

import (
	"context"
	"errors"
	"log"

	"github.com/go-ble/ble"
)

// ReadRSSI returns RSSI of Puck or Pucks
func (p *Puck) ReadRSSI(deviceName ...string) (map[string]int, error) {
	var rssi = make(map[string]int)
	if len(deviceName) > 0 {
		// one or more pucks
		for _, v := range deviceName {
			if _, ok := p.device[v]; !ok {
				return rssi, errors.New("device not found")
			}
			rssi[v] = p.readrssi(p.device[v])
		}
	} else {
		// all pucks
		for k, v := range p.device {
			rssi[k] = p.readrssi(v)
		}
	}
	return rssi, nil
}

// readrssi performs the actual read
func (p *Puck) readrssi(device ble.Addr) int {
	log.Println("reading RSSI...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = ble.WithSigHandler(ctx, func() {
		log.Fatal("cancel signal received...")
	})

	client, err := ble.Dial(ctx, device)
	if err != nil {
		log.Println(err)
	}
	return client.ReadRSSI()
}
