# puckUART

## A go package for working with Espruino Puck.js over Bluetooth LE

Uses the Nordic UART service. Supports both TX (read) and RX (write) characteristics. Includes some Espruino API wrapper methods.

### Usage

Some examples of current usage below.

### Write examples

```go
// Scan for 5 seconds
// accepts optional second param to filter based on substring of name
puck := puckUART.Scan(5 * time.Second)

// Broadcast to all Pucks
puck.Broadcast("LED1.set();")

// Write to single Puck
err = puck.Write("LED1.set();", "Puck.js bb18")

// Use a wrapper func to turn on LED1 on all Pucks
puck.LED1Set()

// As above but specific Puck
puck.LED1Set("Puck.js bb18")

// Make all Pucks load the contents of their flash
// accepts optional name param for specific Puck
puck.Load()

// Make all Pucks run a function loaded from flash
// accepts optional name param for specific Puck
puck.Func("alarm")

// Reset all Pucks
// accepts optional name param for specific Puck
puck.Reset()
```

### Read/Subscribe examples
```go
// Scan for 5 seconds
// accepts optional second param to filter based on substring of name
p := puckUART.Scan(5 * time.Second)

// Obtain RSSI of Puck or Pucks
// accepts optional name param for specific Puck
rssi, _:= puck.ReadRSSI()
for _, v := range rssi {
	log.Println(v)
}

// Subscribe to TX characterstic on all Pucks 
// accepts optional name param for subscribing to specific Puck
p.Subscribe()

// Set up a listener for messages
for msg := range p.Message {
	log.Println(msg.Payload, msg.Device, msg.Timestamp)
}
```

## License

MIT
