package puckUART

import (
	"time"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"

	"context"
	"log"
	"strings"
)

// Puck stores addresses of detected pucks
type Puck struct {
	device  map[string]ble.Addr
	Message chan Message
}

// Message used for reading
type Message struct {
	Payload   string
	Device    string
	Timestamp time.Time
}

var uartServiceID = ble.MustParse("6e400001-b5a3-f393-e0a9-e50e24dcca9e")
var uartServiceRXCharID = ble.MustParse("6e400002-b5a3-f393-e0a9-e50e24dcca9e")
var uartServiceTXCharID = ble.MustParse("6e400003-b5a3-f393-e0a9-e50e24dcca9e")

// Scan looks for pucks and can match optional filter
func Scan(duration time.Duration, filter ...string) *Puck {
	var filterString = "Puck.js"
	p := Puck{}
	p.device = make(map[string]ble.Addr)
	p.Message = make(chan Message)

	// scan filter?
	if len(filter) == 1 {
		filterString = filter[0]
	}

	ftr := func(a ble.Advertisement) bool {
		return strings.Contains(a.LocalName(), filterString)
	}
	// init bt device
	d, err := dev.NewDevice("default")
	if err != nil {
		log.Fatalf("can't init bt device : %s", err)
	}
	ble.SetDefaultDevice(d)

	// scan
	log.Printf("scanning for %s...\n", duration)
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), duration))
	ble.Scan(ctx, true, p.found, ftr)

	return &p
}

// found is a handler for found devices
func (p *Puck) found(a ble.Advertisement) {
	// store address for direct access
	//log.Printf("found %s", a.LocalName())
	_, ok := p.device[a.LocalName()]
	if !ok {
		log.Printf("found %s", a.LocalName())
		p.device[a.LocalName()] = a.Addr()
	}
}
