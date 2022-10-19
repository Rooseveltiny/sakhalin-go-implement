package message

import (
	"sakhalin/color"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func fetchByteArray() []byte {
	return []byte{1, 63, 228, 122, 225, 71, 174, 20, 123, 0, 0, 82, 8, 1, 100, 200, 255, 0, 0, 0, 12, 72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33}
}

func TestMessageInit(t *testing.T) {
	Convey("test message init", t, func() {
		newMessage := NewMessage([]byte{})
		So(len(newMessage.message), ShouldEqual, 0)
		Convey("test add one single byte", func() {
			newMessage.appendSingleByte(0x1)
			So(newMessage.message[0], ShouldEqual, 0x1)
			So(len(newMessage.message), ShouldEqual, 1)
			Convey("test add float64", func() {
				newMessage.appendFloat64(0.64)
				So(len(newMessage.message), ShouldEqual, 9)
				So(newMessage.message[3], ShouldEqual, 0x7a)
				Convey("test add uint32", func() {
					newMessage.appendUint32(21000)
					So(len(newMessage.message), ShouldEqual, 13)
					So(cap(newMessage.message), ShouldEqual, 16)
					So(newMessage.message[10], ShouldEqual, 0x0)
					// two bytes left free because uint32 isn't so much big
					Convey("test add bool", func() {
						newMessage.appendBoolean(true)
						So(newMessage.message[13], ShouldEqual, 0b1)
						Convey("test add bytes", func() {
							newMessage.appendBytes([]byte{100, 200, 255})
							So(newMessage.message[14], ShouldEqual, 0b01100100)
							So(newMessage.message[15], ShouldEqual, 0b11001000)
							So(newMessage.message[16], ShouldEqual, 0b11111111)
							Convey("test add string", func() {
								newMessage.appendString("Hello world!")
								bytes := []byte{
									0x0, 0x0, 0x0, 0xc, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x21,
								}
								So(newMessage.message[17:], ShouldResemble, bytes)
								So(newMessage.message, ShouldResemble, fetchByteArray())
							})
						})
					})
				})
			})
		})
	})
}

func TestAppendColor(t *testing.T) {
	Convey("append color to message", t, func() {
		c := color.NewColor(100, 100, 23, 30)
		m := NewMessage([]byte{})
		m.appendColor(c)
		So(len(m.message), ShouldEqual, 4)
		firstByte := m.retrieveSingleByte()
		So(firstByte, ShouldEqual, 0x64)
	})
}

func TestMessageReading(t *testing.T) {
	Convey("read single byte", t, func() {
		byteArray := fetchByteArray()
		newMessage := NewMessage(byteArray)
		So(len(newMessage.message), ShouldEqual, 33)
		b := newMessage.retrieveSingleByte()
		So(b, ShouldEqual, 0x1)
		So(len(newMessage.message), ShouldEqual, 32)
		Convey("read float64", func() {
			r := newMessage.retrieveFloat64()
			So(r, ShouldEqual, 0.64)
			Convey("read uint32", func() {
				r := newMessage.retrieveUint32()
				So(r, ShouldEqual, 21000)
				Convey("read bool value", func() {
					r := newMessage.retrieveBool()
					So(r, ShouldEqual, true)
					Convey("test left bytes", func() {
						So(len(newMessage.message), ShouldEqual, 19)
					})
					Convey("test read string", func() {
						for i := 0; i < 3; i++ {
							newMessage.retrieveSingleByte()
						}
						s := newMessage.retrieveString()
						So(s, ShouldEqual, "Hello world!")
					})
				})
			})
		})
	})
}
