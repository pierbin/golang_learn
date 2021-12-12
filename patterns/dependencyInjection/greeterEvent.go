package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

//In Wire, initializers are known as "providers," functions which provide a particular type.

type Message string

func NewMessage(phrase string) Message {
	return Message(phrase)
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

type Greeter struct {
	Message Message // <- adding a Message field
	Grumpy  bool
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

//In Wire, initializers are known as "providers," functions which provide a particular type.
//We add a zero value for Event as a return value to satisfy the compiler.
//Note that even if we add values to Event, Wire will ignore them.
//In fact, the injector's purpose is to provide information about which providers to use to construct an Event.

//First create a message,
//then we create a greeter with that message,
//and finally we create an event with that greeter.
//With all the initialization done, we're ready to start our event.

func main() {
	e, err := InitializeEvent("")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
