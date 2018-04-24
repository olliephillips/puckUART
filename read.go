package puckUART

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/go-ble/ble"
)

// Subscribe starts listening to TX on a single Puck or multiple Pucks
func (p *Puck) Subscribe(deviceName ...string) error {
	var err error
	var check = func(name string) error {
		if _, ok := p.device[name]; !ok {
			return errors.New(name + " not found")
		}
		return nil
	}

	if len(deviceName) == 1 {
		// subscribe single puck
		err = check(deviceName[0])
		if err != nil {
			return err
		}
		go p.subscribe(deviceName[0], p.device[deviceName[0]])
	} else {
		// subscribe all pucks
		for k, v := range p.device {
			go p.subscribe(k, v)
		}
	}
	return nil
}

// sets up the subscription to TX for single Puck
func (p *Puck) subscribe(deviceName string, device ble.Addr) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = ble.WithSigHandler(ctx, func() {
		log.Fatal("cancel signal received")
	})

	client, _ := ble.Dial(ctx, device)
	services, _ := client.DiscoverServices(nil)
	for _, s := range services {
		if s.UUID.Equal(uartServiceID) {
			characteristics, _ := client.DiscoverCharacteristics(nil, s)
			for _, c := range characteristics {
				if c.UUID.Equal(uartServiceTXCharID) {
					err := client.Subscribe(c, false, func(pl []byte) {
						// cleanup
						data := strings.Replace(string(pl), ">", "", -1)
						data = strings.Replace(data, "\r\n", "", -1)
						// assign to struct
						msg := Message{}
						msg.Payload = data
						msg.Device = deviceName
						msg.Timestamp = time.Now()
						// put on channel
						p.Message <- msg
					})
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}
