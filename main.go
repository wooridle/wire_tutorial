package main

import (
	"fmt"
	"os"
	"wire_tutorial/wire_tutorial"
)

func main()  {
	//beforeWire()
	//NewMessage("hi there")

	//greeter := NewGreeter(NewMessage("hi there"))
	//e := NewEvent(greeter)
	//fmt.Println(e.Greeter.Greet())
	//newEvent := NewEvent(NewGreeter(NewMessage("hi there")))
	//newEvent.Start()

	event, err := wire_tutorial.InitializeEvent()
	if err != nil {
		fmt.Printf("failed to create event: %s\n\n", err)
		os.Exit(2)
	}
	event.Start()
}
