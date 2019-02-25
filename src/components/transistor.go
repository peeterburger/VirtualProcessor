package main

import "fmt"

type Transistor struct {
	collector *Pin
	base      *Pin
	emitter   *Pin
}

func NewTransistor() *Transistor {
	t := Transistor{}
	t.collector = &Pin{root: &t}
	t.base = &Pin{root: &t}
	t.emitter = &Pin{root: &t}
	return &t
}

func main() {
	t1 := NewTransistor()
	t2 := NewTransistor()

	t1.emitter.ConnectTo(t2.emitter)
	
	t1.collector.POWER_SUPPLY()
	t2.collector.POWER_SUPPLY()

	t1.base.INPUT(false)
	t2.base.INPUT(false)

	fmt.Println(t2.emitter.OUTPUT())
}
