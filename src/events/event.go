package events

import "sakhalin/flags"

type Event interface {
	bitmask() flags.EventMask
}
