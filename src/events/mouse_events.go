package events

import "sakhalin/flags"

type MouseEvent struct {
	Buttons flags.MouseButtons
	X       int
	Y       int
	Mod     flags.ModifierKeys
}

func (m MouseEvent) Bitmask() flags.EventMask {
	return flags.I_maskMouseMove | flags.I_maskMouseUp | flags.I_maskMouseDown | flags.I_maskClick | flags.I_maskDblClick | flags.I_maskAuxClick
}

type MouseMoveEvent struct{ MouseEvent }

func (m MouseMoveEvent) Bitmask() flags.EventMask { return flags.I_maskMouseMove }

type MouseDownEvent struct{ MouseEvent }

func (e MouseDownEvent) Bitmask() flags.EventMask { return flags.I_maskMouseDown }

type MouseUpEvent struct{ MouseEvent }

func (e MouseUpEvent) Bitmask() flags.EventMask { return flags.I_maskMouseUp }

type MouseClickEvent struct{ MouseEvent }

func (e MouseClickEvent) Bitmask() flags.EventMask { return flags.I_maskClick }

type MouseDoubleClickEvent struct{ MouseEvent }

func (e MouseDoubleClickEvent) Bitmask() flags.EventMask { return flags.I_maskDblClick }

type MouseAuxClickEvent struct{ MouseEvent }

func (e MouseAuxClickEvent) Bitmask() flags.EventMask { return flags.I_maskAuxClick }

type WheelEvent struct {
	MouseEvent
	DeltaX    float64
	DeltaY    float64
	DeltaZ    float64
	DeltaMode flags.DeltaMode
}

func (e WheelEvent) Bitmask() flags.EventMask { return flags.I_maskWheel }
