// go:build wireinject
//+build wireinject

package wire_tutorial

import "github.com/google/wire"

func InitializeEvent() (Event, error) {
	wire.Build(NewEvent2, NewGreeter2, NewMessage)
	return Event{}, nil
}
