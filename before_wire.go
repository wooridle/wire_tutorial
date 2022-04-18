package wire_tutorial

import (
	"errors"
	"fmt"
	"time"
)

type Message string

type Greeter struct {
	Message Message
	Grumpy  bool
}

type Event struct {
	Greeter Greeter
}

func NewMessage() Message {
	return Message("hi there")
}

//func NewGreeter(m Message) Greeter {
//	return Greeter{Message: m}
//}

func NewGreeter2(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}


//func (g Greeter) Greet() Message {
//	return g.Message
//}

func (g Greeter) Greet2() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

//func NewEvent(g Greeter) Event {
//	return Event{Greeter: g}
//}

func NewEvent2(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start()  {
	msg := e.Greeter.Greet2()
	fmt.Println(msg)
}

//func beforeWire()  {
//	message := NewMessage("before wire / hi there")
//	greeter := NewGreeter(message)
//	event := NewEvent(greeter)
//
//	event.Start()
//}
//
//func wireImprove()  {
//	message := NewMessage("before wire / hi there")
//	greeter := NewGreeter(message)
//	event := NewEvent(greeter)
//
//	event.Start()
//}
