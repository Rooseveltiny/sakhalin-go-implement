package flags

/*
	Delta mode allows us to retrieve data about changing of mouse scrolling
*/
type DeltaMode byte

const (
	B_DeltaPixel DeltaMode = 0b00000001
	B_DeltaLine  DeltaMode = 0b00000010
	B_DeltaPage  DeltaMode = 0b00000100
)

/*
	Mouse modifier keys helps to know whether special buttons are depressed while event
*/

type ModifierKeys byte

const (
	B_modKeyAlt   ModifierKeys = 0b00000001
	B_modKeyShift ModifierKeys = 0b00000010
	B_modKeyCtrl  ModifierKeys = 0b00000100
	B_modKeyMeta  ModifierKeys = 0b00001000
)

/*
	Event mask is a special parameters to know which buttons are depressed simultaneously
*/

type EventMask int

const (
	I_maskMouseMove   EventMask = 1
	I_maskMouseDown   EventMask = 2
	I_maskMouseUp     EventMask = 4
	I_maskKeyDown     EventMask = 8
	I_maskKeyUp       EventMask = 16
	I_maskClick       EventMask = 32
	I_maskDblClick    EventMask = 64
	I_maskAuxClick    EventMask = 128
	I_maskWheel       EventMask = 256
	I_maskTouchStart  EventMask = 512
	I_maskTouchMove   EventMask = 1024
	I_maskTouchEnd    EventMask = 2048
	I_maskTouchCancel EventMask = 4096
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
