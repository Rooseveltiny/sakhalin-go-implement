package events

import "sakhalin/flags"

type Event interface {
	Bitmask() flags.EventMask
}
