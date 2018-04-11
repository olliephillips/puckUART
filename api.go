package puckUART

// Reset resets one or more pucks
func (p *Puck) Reset(name ...string) error {
	cmd := []byte("reset();\n")
	err := p.command(name, cmd)
	if err != nil {
		return err
	}
	return nil
}

// Func will call a saved function (load from flash first)
func (p *Puck) Func(function string, name ...string) error {
	cmd := []byte(function + "();\n")
	err := p.command(name, cmd)
	if err != nil {
		return err
	}
	return nil
}

// Load makes ready saved code from flash
func (p *Puck) Load(name ...string) error {
	cmd := []byte("load();\n")
	err := p.command(name, cmd)
	if err != nil {
		return err
	}
	return nil
}

// LED1Set turns on LED1
func (p *Puck) LED1Set(name ...string) error {
	cmd := []byte("LED1.set();\n")
	err := p.command(name, cmd)
	if err != nil {
		return err
	}
	return nil
}

// LED1Reset turns off LED1
func (p *Puck) LED1Reset(name ...string) error {
	cmd := []byte("LED1.reset();\n")
	err := p.command(name, cmd)
	if err != nil {
		return err
	}
	return nil
}

// LED2Set turns on LED2
func (p *Puck) LED2Set(name ...string) error {
	cmd := []byte("LED2.set();\n")
	err := p.command(name, cmd)
	if err != nil {
		return err
	}
	return nil
}

// LED2Reset turns off LED2
func (p *Puck) LED2Reset(name ...string) error {
	cmd := []byte("LED2.reset();\n")
	err := p.command(name, cmd)
	if err != nil {
		return err
	}
	return nil
}

// LED3Set turns on LED3
func (p *Puck) LED3Set(name ...string) error {
	cmd := []byte("LED3.set();\n")
	err := p.command(name, cmd)
	if err != nil {
		return err
	}
	return nil
}

// LED3Reset turns off LED3
func (p *Puck) LED3Reset(name ...string) error {
	cmd := []byte("LED3.reset();\n")
	err := p.command(name, cmd)
	if err != nil {
		return err
	}
	return nil
}
