package flags

/*
	Delta mode allows us to retrieve data about changing of mouse scrolling
*/
type DeltaMode byte

const (
	bDeltaPixel DeltaMode = iota
	bDeltaLine
	bDeltaPage
)

/*
	Mouse modifier keys helps to know whether special buttons are depressed while event
*/

type ModifierKeys byte

const (
	bmodKeyAlt ModifierKeys = 1 << iota
	bmodKeyShift
	bmodKeyCtrl
	bmodKeyMeta
)

/*
	Event mask is a special parameters to know which buttons are depressed simultaneously
*/

type eventMask int

const (
	imaskMouseMove eventMask = 1 << iota
	imaskMouseDown
	imaskMouseUp
	imaskKeyDown
	imaskKeyUp
	imaskClick
	imaskDblClick
	imaskAuxClick
	imaskWheel
	imaskTouchStart
	imaskTouchMove
	imaskTouchEnd
	imaskTouchCancel
)

/*
	Mouse buttons enum represented in bytes
*/

type MouseButtons byte

const (
	ButtonPrimary MouseButtons = 1 << iota
	ButtonSecondary
	ButtonAuxiliary
	Button4th
	Button5th
	ButtonNone MouseButtons = 0
)

/*
	General mouse events
*/

type MouseEvent byte

const (
	bMouseMove MouseEvent = 1 + iota
	bMouseDown
	bMouseUp
	bKeyDown
	bKeyUp
	bClick
	bDblClick
	bAuxClick
	bWheel
	bTouchStart
	bTouchMove
	bTouchEnd
	bTouchCancel
)
