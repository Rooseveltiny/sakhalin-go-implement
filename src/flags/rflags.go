package flags

/*
	Delta mode allows us to retrieve data about changing of mouse scrolling
*/
type DeltaMode byte

const (
	B_DeltaPixel DeltaMode = 0x0
	B_DeltaLine  DeltaMode = 0x1
	B_DeltaPage  DeltaMode = 0x2
)

/*
	Mouse modifier keys helps to know whether special buttons are depressed while event
*/

type ModifierKeys byte

const (
	B_modKeyAlt   ModifierKeys = 0x1
	B_modKeyShift ModifierKeys = 0x2
	B_modKeyCtrl  ModifierKeys = 0x4
	B_modKeyMeta  ModifierKeys = 0x8
)

/*
	Event mask is a special parameters to know which buttons are depressed simultaneously
*/

type EventMask int

const (
	I_maskMouseMove   EventMask = 0b0000000000000001
	I_maskMouseDown   EventMask = 0b0000000000000010
	I_maskMouseUp     EventMask = 0b0000000000000100
	I_maskKeyDown     EventMask = 0b0000000000001000
	I_maskKeyUp       EventMask = 0b0000000000010000
	I_maskClick       EventMask = 0b0000000000100000
	I_maskDblClick    EventMask = 0b0000000001000000
	I_maskAuxClick    EventMask = 0b0000000010000000
	I_maskWheel       EventMask = 0b0000000100000000
	I_maskTouchStart  EventMask = 0b0000001000000000
	I_maskTouchMove   EventMask = 0b0000010000000000
	I_maskTouchEnd    EventMask = 0b0000100000000000
	I_maskTouchCancel EventMask = 0b0001000000000000
)

/*
	Mouse buttons enum represented in bytes
*/

type MouseButtons byte

const (
	B_ButtonPrimary   MouseButtons = 0b00000001
	B_ButtonSecondary MouseButtons = 0b00000010
	B_ButtonAuxiliary MouseButtons = 0b00000100
	B_Button4th       MouseButtons = 0b00001000
	B_Button5th       MouseButtons = 0b00010000
	B_ButtonNone      MouseButtons = 0b00000000
)

/*
	General mouse events
*/

type MouseEvent byte

const (
	B_MouseMove   MouseEvent = 1
	B_MouseDown   MouseEvent = 2
	B_MouseUp     MouseEvent = 3
	B_KeyDown     MouseEvent = 4
	B_KeyUp       MouseEvent = 5
	B_Click       MouseEvent = 6
	B_DblClick    MouseEvent = 7
	B_AuxClick    MouseEvent = 8
	B_Wheel       MouseEvent = 9
	B_TouchStart  MouseEvent = 10
	B_TouchMove   MouseEvent = 11
	B_TouchEnd    MouseEvent = 12
	B_TouchCancel MouseEvent = 13
)
