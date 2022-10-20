package events

import (
	"sakhalin/message"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func retrieveMessageMouseMove() *message.Message {
	return message.NewMessage([]byte{1, 0b00000001, 0x0, 0x0, 0x0, 200, 0x0, 0x0, 0x0, 200, 0x4})
}

func retrieveMessageMouseDown() *message.Message {
	return message.NewMessage([]byte{2, 0b00000000, 0x0, 0x0, 0x0, 255, 0x0, 0x0, 0x1, 0xAF, 0b0001000})
}

func retrieveMessageMouseUp() *message.Message {
	return message.NewMessage([]byte{3, 0b00000001, 0x0, 0x0, 0x11, 0x95, 0x0, 0x0, 0x2, 0xBC, 0b0001000})
}

func retrieveMessageKeyDown() *message.Message {
	return message.NewMessage([]byte{4, 0b00000010, 0x0, 0x0, 0x0, 0x1, 0x68})
}

func retrieveMessageKeyUp() *message.Message {
	return message.NewMessage([]byte{5, 0b00000100, 0x0, 0x0, 0x0, 0x1, 0x59})
}

func retrieveMessageMouseClick() *message.Message {
	return message.NewMessage([]byte{6, 0b00000001, 0x0, 0x0, 0x11, 0x95, 0x0, 0x0, 0x2, 0xBC, 0b0001000})
}

func retrieveMessageDoubleMouseClick() *message.Message {
	return message.NewMessage([]byte{7, 0b00000001, 0x0, 0x0, 0x11, 0x95, 0x0, 0x0, 0x2, 0xBC, 0b0001000})
}

func retrieveMessageMouseAuxClick() *message.Message {
	return message.NewMessage([]byte{9, 0b00000001, 0x0, 0x0, 0x11, 0x95, 0x0, 0x0, 0x2, 0xBC, 0b0001000})
}

func retrieveMessageMouseWheel() *message.Message {
	wheelEvent := byte(9)
	mouseEvent := []byte{0b00000000, 0x0, 0x0, 0x6b, 0x89, 0x0, 0x0, 0x0, 0x1, 0b00000010}
	DeltaX := []byte{0x40, 0x74, 0x40, 0x90, 0x38, 0x5c, 0x67, 0xe0}
	DeltaY := []byte{0x40, 0x81, 0xa3, 0x4c, 0xec, 0x41, 0xdd, 0x1a}
	DeltaZ := []byte{0x40, 0x74, 0x40, 0x90, 0x38, 0x5c, 0x67, 0xe0}
	deltaMode := 0x0
	bArray := make([]byte, 0)
	bArray = append(bArray, wheelEvent)
	bArray = append(bArray, mouseEvent...)
	bArray = append(bArray, DeltaX...)
	bArray = append(bArray, DeltaY...)
	bArray = append(bArray, DeltaZ...)
	bArray = append(bArray, byte(deltaMode))

	return message.NewMessage(bArray)
}

func TestEventDecoding(t *testing.T) {
	Convey("check mouse move event", t, func() {
		m := retrieveMessageMouseMove()
		ev, _ := DecodeEvent(m)
		switch e := ev.(type) {
		case MouseMoveEvent:
			So(e.X, ShouldEqual, 200)
			So(e.Y, ShouldEqual, 200)
			So(e.Mod, ShouldEqual, 0x4)
			So(e.Buttons, ShouldEqual, 0b00000001)
		}
	})
	Convey("check mouse down event", t, func() {
		m := retrieveMessageMouseDown()
		ev, _ := DecodeEvent(m)
		switch e := ev.(type) {
		case MouseDownEvent:
			So(e.X, ShouldEqual, 255)
			So(e.Y, ShouldEqual, 431)
			So(e.Mod, ShouldEqual, 0b00001000)
			So(e.Buttons, ShouldEqual, 0b00000000)
		}
	})
	Convey("check mouse up event", t, func() {
		m := retrieveMessageMouseUp()
		ev, _ := DecodeEvent(m)
		switch e := ev.(type) {
		case MouseUpEvent:
			So(e.X, ShouldEqual, 4501)
			So(e.Y, ShouldEqual, 700)
			So(e.Mod, ShouldEqual, 0b00001000)
			So(e.Buttons, ShouldEqual, 0b00000001)
		}
	})
	Convey("check mouse click event", t, func() {
		m := retrieveMessageMouseClick()
		ev, _ := DecodeEvent(m)
		switch e := ev.(type) {
		case MouseClickEvent:
			So(e.X, ShouldEqual, 4501)
			So(e.Y, ShouldEqual, 700)
			So(e.Mod, ShouldEqual, 0b00001000)
			So(e.Buttons, ShouldEqual, 0b00000001)
		}
	})
	Convey("check mouse double click event", t, func() {
		m := retrieveMessageDoubleMouseClick()
		ev, _ := DecodeEvent(m)
		switch e := ev.(type) {
		case MouseDoubleClickEvent:
			So(e.X, ShouldEqual, 4501)
			So(e.Y, ShouldEqual, 700)
			So(e.Mod, ShouldEqual, 0b00001000)
			So(e.Buttons, ShouldEqual, 0b00000001)
		}
	})
	Convey("check mouse aux click event", t, func() {
		m := retrieveMessageMouseAuxClick()
		ev, _ := DecodeEvent(m)
		switch e := ev.(type) {
		case MouseAuxClickEvent:
			So(e.X, ShouldEqual, 4501)
			So(e.Y, ShouldEqual, 700)
			So(e.Mod, ShouldEqual, 0b00001000)
			So(e.Buttons, ShouldEqual, 0b00000001)
		}
	})
	Convey("check mouse wheel event", t, func() {
		m := retrieveMessageMouseWheel()
		ev, _ := DecodeEvent(m)
		switch e := ev.(type) {
		case MouseWheelEvent:
			So(e.MouseEvent.X, ShouldEqual, 27529)
			So(e.DeltaX, ShouldEqual, 324.03521)
			So(e.DeltaY, ShouldEqual, 564.41256)
			So(e.DeltaZ, ShouldEqual, 324.03521)
			So(e.DeltaMode, ShouldEqual, 0x0)
		}
	})
	Convey("check key down event", t, func() {
		m := retrieveMessageKeyDown()
		ev, _ := DecodeEvent(m)
		switch e := ev.(type) {
		case KeyDownEvent:
			So(e.Key, ShouldEqual, "h")
			So(e.Mod, ShouldEqual, 0b00000010)
		}
	})
	Convey("check key up event", t, func() {
		m := retrieveMessageKeyUp()
		ev, _ := DecodeEvent(m)
		switch e := ev.(type) {
		case KeyUpEvent:
			So(e.Key, ShouldEqual, "Y")
			So(e.Mod, ShouldEqual, 0b00000100)
		}
	})
}
