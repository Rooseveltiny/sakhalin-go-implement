package events

import (
	"fmt"
	"sakhalin/flags"
	"sakhalin/message"
)

func decodeMouseEvent(m *message.Message) MouseEvent {
	return MouseEvent{
		Buttons: flags.MouseButtons(m.RetrieveSingleByte()),
		X:       int(m.RetrieveUint32()),
		Y:       int(m.RetrieveUint32()),
		Mod:     flags.ModifierKeys(m.RetrieveSingleByte()),
	}
}

func decodeKeyboardEvent(m *message.Message) KeyboardEvent {
	return KeyboardEvent{
		Mod: flags.ModifierKeys(m.RetrieveSingleByte()),
		Key: m.RetrieveString(),
	}
}

func decodeWheelEvent(m *message.Message) WheelEvent {
	return WheelEvent{
		MouseEvent: decodeMouseEvent(m),
		DeltaX:     m.RetrieveFloat64(),
		DeltaY:     m.RetrieveFloat64(),
		DeltaZ:     m.RetrieveFloat64(),
		DeltaMode:  flags.DeltaMode(m.RetrieveSingleByte()),
	}
}

type decodeFunc = func(m *message.Message) Event

var decodePerformers = map[flags.EventMask]decodeFunc{
	flags.I_maskMouseMove: func(m *message.Message) Event { return MouseMoveEvent{decodeMouseEvent(m)} },
	flags.I_maskMouseDown: func(m *message.Message) Event { return MouseDownEvent{decodeMouseEvent(m)} },
	flags.I_maskMouseUp:   func(m *message.Message) Event { return MouseUpEvent{decodeMouseEvent(m)} },
	flags.I_maskKeyDown:   func(m *message.Message) Event { return KeyDownEvent{decodeKeyboardEvent(m)} },
	flags.I_maskKeyUp:     func(m *message.Message) Event { return KeyUpEvent{decodeKeyboardEvent(m)} },
	flags.I_maskClick:     func(m *message.Message) Event { return MouseClickEvent{decodeMouseEvent(m)} },
	flags.I_maskDblClick:  func(m *message.Message) Event { return MouseDoubleClickEvent{decodeMouseEvent(m)} },
	flags.I_maskAuxClick:  func(m *message.Message) Event { return MouseAuxClickEvent{decodeMouseEvent(m)} },
	flags.I_maskWheel:     func(m *message.Message) Event { return MouseWheelEvent{decodeWheelEvent(m)} },
}

func decodeEventFromMessage(m *message.Message) (Event, error) {
	eventType := m.RetrieveSingleByte()
	if decodePerformer, ok := decodePerformers[flags.EventMask(eventType)]; ok {
		return decodePerformer(m), nil
	} else {
		return nil, errUnknownEventType{unknownType: eventType}
	}
}

type errUnknownEventType struct {
	unknownType byte
}

func (err errUnknownEventType) Error() string {
	return fmt.Sprintf("unknown event type: %#x", err.unknownType)
}

func DecodeEvent(m *message.Message) (Event, error) {
	return decodeEventFromMessage(m)
}
