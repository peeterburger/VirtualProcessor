package main

import "fmt"

// Transistor is a struct containing pointers to three pins: Base, Collector and
// Emitter. It represents a physical Transistor and therefore requires some kind
// of power supply and a binary input on the Base-Pin to function.
type Transistor struct {
	Collector *Pin
	Base      *Pin
	Emitter   *Pin
}

// New returns a pointer to a newly created Transistor with default
// initialization values.
func New() *Transistor {
	t := Transistor{}
	t.Collector = &Pin{root: &t}
	t.Base = &Pin{root: &t}
	t.Emitter = &Pin{root: &t}
	return &t
}

func main() {
	t1 := New()
	t2 := New()

	t1.Emitter.ConnectTo(t2.Collector)
	t1.Collector.PowerSupply()

	t1.Base.Input(true)
	t2.Base.Input(true)

	fmt.Println(t1.Collector.Output())
}
