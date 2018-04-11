package puckUART

import (
	"context"
	"errors"
	"log"

	"github.com/go-ble/ble"
)

// Broadcast sends payload to all known pucks
func (p *Puck) Broadcast(payload string) {
	cmd := []byte(payload + "\n")
	for _, a := range p.device {
		p.write(cmd, a)
	}
}

// Write is the public wrapper for write with some error handling
func (p *Puck) Write(payload string, deviceName string) error {
	cmd := []byte(payload + "\n")
	if _, ok := p.device[deviceName]; !ok {
		return errors.New("device not found")

	}
	p.write(cmd, p.device[deviceName])
	return nil
}

func (p *Puck) command(name []string, cmd []byte) error {
	if len(name) > 0 {
		// one or more pucks
		for _, n := range name {
			if _, ok := p.device[n]; !ok {
				return errors.New("cannot '" + string(cmd) + "' device not found")
			}
			p.write(cmd, p.device[n])
		}
	} else {
		// all pucks
		for _, a := range p.device {
			p.write(cmd, a)
		}
	}
	return nil
}

// write performs the writes
func (p *Puck) write(cmd []byte, device ble.Addr) {
	ctx := context.Background()
	client, _ := ble.Dial(ctx, device)
	services, _ := client.DiscoverServices(nil)
	for _, s := range services {
		if s.UUID.Equal(uartServiceID) {
			characteristics, _ := client.DiscoverCharacteristics(nil, s)
			for _, c := range characteristics {
				if c.UUID.Equal(uartServiceRXCharID) {
					log.Printf("writing...%s", string(cmd))
					err := client.WriteCharacteristic(c, cmd, false)
					if err != nil {
						log.Fatalf("can't write : %s", err)
					}
				}
			}
		}
	}
	// close connection
	client.CancelConnection()
}
