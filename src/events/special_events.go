package events

import "sakhalin/flags"

type StopEvent struct{}

func (e StopEvent) bitmask() flags.EventMask { return 0b00000000 }
