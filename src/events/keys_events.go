package events

import "sakhalin/flags"

type KeyboardEvent struct {
	Key string
	Mod flags.ModifierKeys
}

func (e KeyboardEvent) Bitmask() flags.EventMask { return flags.I_maskKeyDown | flags.I_maskKeyUp }

type KeyDownEvent struct{ KeyboardEvent }

func (e KeyDownEvent) Bitmask() flags.EventMask { return flags.I_maskKeyDown }

type KeyUpEvent struct{ KeyboardEvent }

func (e KeyUpEvent) Bitmask() flags.EventMask { return flags.I_maskKeyUp }
