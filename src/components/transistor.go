package main

import "fmt"

// Transistor is a struct containing pointers to three pins: Base, Collector and
// Emitter. It represents a physical Transistor and therefore requires some kind
// of power supply and a binary input on the Base-Pin to function.
type Transistor struct {
	collector *Pin
	base      *Pin
	emitter   *Pin
}

// New returns a pointer to a newly created Transistor with default
// initialization values.
func New() *Transistor {
	t := Transistor{}
	t.collector = &Pin{root: &t}
	t.base = &Pin{root: &t}
	t.emitter = &Pin{root: &t}
	return &t
}

func main() {
	t1 := New()
	t2 := New()

	t1.emitter.ConnectTo(t2.emitter)

	t1.collector.PowerSupply()
	t2.collector.PowerSupply()

	t1.base.Input(true)
	t2.base.Input(true)

	fmt.Println(t2.emitter.Output())
}
