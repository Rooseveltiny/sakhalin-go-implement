package flags

/*
	Delta mode allows us to retrieve data about changing of mouse scrolling
*/
type DeltaMode byte

const (
	B_DeltaPixel DeltaMode = iota
	B_DeltaLine
	B_DeltaPage
)

/*
	Mouse modifier keys helps to know whether special buttons are depressed while event
*/

type ModifierKeys byte

const (
	B_modKeyAlt ModifierKeys = 1 << iota
	B_modKeyShift
	B_modKeyCtrl
	B_modKeyMeta
)

/*
	Event mask is a special parameters to know which buttons are depressed simultaneously
*/

type EventMask int

const (
	I_maskMouseMove EventMask = 1 << iota
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
	B_ButtonPrimary MouseButtons = 1 << iota
	B_ButtonSecondary
	B_ButtonAuxiliary
	B_Button4th
	B_Button5th
	B_ButtonNone MouseButtons = 0
)

/*
	General mouse events
*/

type MouseEvent byte

const (
	B_MouseMove MouseEvent = 1 + iota
	B_MouseDown
	B_MouseUp
	B_KeyDown
	B_KeyUp
	B_Click
	B_DblClick
	B_AuxClick
	B_Wheel
	B_TouchStart
	B_TouchMove
	B_TouchEnd
	B_TouchCancel
)
