package main

type Pin struct {
	connected []*Pin
	root      *Transistor
	state     uint8
}

const (
	STATE_FALSE        uint8 = 0
	STATE_TRUE         uint8 = 1
	STATE_POWER_SUPPLY uint8 = 2
	STATE_MASS         uint8 = 3
)

func (p *Pin) ConnectTo(remotePin *Pin) {
	p.connected = append(p.connected, remotePin)
	remotePin.connected = append(remotePin.connected, p)
}

func (p *Pin) INPUT(input_state bool) {
	if input_state {
		p.state = STATE_TRUE
	}
}

func (p *Pin) POWER_SUPPLY() {
	p.state = STATE_POWER_SUPPLY
}

func (p *Pin) MASS() {
	p.state = STATE_MASS
}

func (p *Pin) OUTPUT() bool {
	return p.recOutput(nil)
}

func (p *Pin) recOutput(sourcePin *Pin) bool {
	// fetch and evaluate pin state
	switch p.state {
	case STATE_FALSE:
		break
	case STATE_TRUE:
		return true
	case STATE_POWER_SUPPLY:
		return true
	case STATE_MASS:
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
