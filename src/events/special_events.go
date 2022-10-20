package events

import "sakhalin/flags"

type StopEvent struct{}

func (e StopEvent) Bitmask() flags.EventMask { return 0b00000000 }
