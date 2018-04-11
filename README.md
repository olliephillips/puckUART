# puckUART

## A go package for working with Espruino Puck.js over Bluetooth LE

Uses the Nordic UART service. Currently only supports RX characteristic as sufficient for my project.

Includes some Espruino API wrapper methods. Hope to add some TX methods to read console and status

### Usage

Some examples of current usage below:

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
// accepts optional name param for specific puck
puck.Reset()

```

## License

MIT
