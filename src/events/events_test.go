package events

import (
	"sakhalin/message"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func retrieveMessageMouseMove() *message.Message {
	mouseMoveBytes := []byte{0b0000000000000001, 0b00000001, 0x0, 0x0, 0x0, 200, 0x0, 0x0, 0x0, 200, 0x4}
	return message.NewMessage(mouseMoveBytes)
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
}
