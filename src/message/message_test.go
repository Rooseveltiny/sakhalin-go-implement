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
		So(len(newMessage.Message), ShouldEqual, 0)
		Convey("test add one single byte", func() {
			newMessage.AppendSingleByte(0x1)
			So(newMessage.Message[0], ShouldEqual, 0x1)
			So(len(newMessage.Message), ShouldEqual, 1)
			Convey("test add float64", func() {
				newMessage.AppendFloat64(0.64)
				So(len(newMessage.Message), ShouldEqual, 9)
				So(newMessage.Message[3], ShouldEqual, 0x7a)
				Convey("test add uint32", func() {
					newMessage.AppendUint32(21000)
					So(len(newMessage.Message), ShouldEqual, 13)
					So(cap(newMessage.Message), ShouldEqual, 16)
					So(newMessage.Message[10], ShouldEqual, 0x0)
					// two bytes left free because uint32 isn't so much big
					Convey("test add bool", func() {
						newMessage.AppendBoolean(true)
						So(newMessage.Message[13], ShouldEqual, 0b1)
						Convey("test add bytes", func() {
							newMessage.AppendBytes([]byte{100, 200, 255})
							So(newMessage.Message[14], ShouldEqual, 0b01100100)
							So(newMessage.Message[15], ShouldEqual, 0b11001000)
							So(newMessage.Message[16], ShouldEqual, 0b11111111)
							Convey("test add string", func() {
								newMessage.AppendString("Hello world!")
								bytes := []byte{
									0x0, 0x0, 0x0, 0xc, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x21,
								}
								So(newMessage.Message[17:], ShouldResemble, bytes)
								So(newMessage.Message, ShouldResemble, fetchByteArray())
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
		m.AppendColor(c)
		So(len(m.Message), ShouldEqual, 4)
		firstByte := m.RetrieveSingleByte()
		So(firstByte, ShouldEqual, 0x64)
	})
}

func TestMessageReading(t *testing.T) {
	Convey("read single byte", t, func() {
		byteArray := fetchByteArray()
		newMessage := NewMessage(byteArray)
		So(len(newMessage.Message), ShouldEqual, 33)
		b := newMessage.RetrieveSingleByte()
		So(b, ShouldEqual, 0x1)
		So(len(newMessage.Message), ShouldEqual, 32)
		Convey("read float64", func() {
			r := newMessage.RetrieveFloat64()
			So(r, ShouldEqual, 0.64)
			Convey("read uint32", func() {
				r := newMessage.RetrieveUint32()
				So(r, ShouldEqual, 21000)
				Convey("read bool value", func() {
					r := newMessage.RetrieveBool()
					So(r, ShouldEqual, true)
					Convey("test left bytes", func() {
						So(len(newMessage.Message), ShouldEqual, 19)
					})
					Convey("test read string", func() {
						for i := 0; i < 3; i++ {
							newMessage.RetrieveSingleByte()
						}
						s := newMessage.RetrieveString()
						So(s, ShouldEqual, "Hello world!")
					})
				})
			})
		})
	})
}

func TestClearMessagePool(t *testing.T) {
	Convey("test clear func", t, func() {
		byteArray := fetchByteArray()
		newMessage := NewMessage(byteArray)
		newMessage.Clear()
		So(len(newMessage.Message), ShouldEqual, 0)
	})
}
