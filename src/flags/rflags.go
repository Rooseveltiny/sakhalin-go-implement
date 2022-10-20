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
	I_maskMouseMove EventMask = 1 + iota
	I_maskMouseDown
	I_maskMouseUp
	I_maskKeyDown
	I_maskKeyUp
	I_maskClick
	I_maskDblClick
	I_maskAuxClick
	I_maskWheel
	I_maskTouchStart
	I_maskTouchMove
	I_maskTouchEnd
	I_maskTouchCancel
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
