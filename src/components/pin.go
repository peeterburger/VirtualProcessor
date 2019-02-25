package main

// Pin is a struct containing an array with pointers to other connected pins, a
// pointer to the pins root/parent element and a specivic state
type Pin struct {
	connected []*Pin
	root      *Transistor
	state     uint8
}

const (
	// StateFalse -> 0
	StateFalse uint8 = 0
	// StateTrue -> 1
	StateTrue uint8 = 1
	// StatePowerSupply -> 2
	StatePowerSupply uint8 = 2
	// StateMass -> 3
	StateMass uint8 = 3
)

// ConnectTo connects two pins
func (p *Pin) ConnectTo(remotePin *Pin) {
	p.connected = append(p.connected, remotePin)
	remotePin.connected = append(remotePin.connected, p)
}

// Input sets the current state of the pin to true/false
func (p *Pin) Input(inputState bool) {
	if inputState {
		p.state = StateTrue
	}
}

// PowerSupply sets the current state of the pin to PowerSupply
func (p *Pin) PowerSupply() {
	p.state = StatePowerSupply
}

// PowerSupply sets the current state of the pin to Mass
func (p *Pin) Mass() {
	p.state = StateMass
}

// Output recursively calculates the output on the current pin
func (p *Pin) Output() bool {
	return p.recOutput(nil)
}

func (p *Pin) recOutput(sourcePin *Pin) bool {
	// fetch and evaluate pin state
	switch p.state {
	case StateFalse:
		break
	case StateTrue:
		return true
	case StatePowerSupply:
		return true
	case StateMass:
		return false
	}

	// fetch remote pins
	var remotePins []*Pin
	for _, rp := range p.connected {
		if rp != p && rp != sourcePin {
			remotePins = append(remotePins, rp)
		}
	}

	// fetch root
	root := p.root
	if root == nil {
		for _, rp := range remotePins {
			if rp.recOutput(p) {
				return true
			}
		}
		return false
	}

	// determines emitter output
	if p == root.emitter {
		if root.collector.recOutput(p) && root.base.recOutput(p) {
			return true
		}
	}

	// OR
	for _, rp := range remotePins {
		if rp.recOutput(p) {
			return true
		}
	}
	return false
}
