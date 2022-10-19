package message

import (
	"encoding/binary"
	"math"
	"sakhalin/color"
)

type Message struct {
	message []byte
	err     error
}

func NewMessage(b []byte) *Message {
	return &Message{
		message: b,
		err:     nil,
	}
}

func (m *Message) appendSingleByte(b byte) {
	m.message = append(m.message, b)
}

func (m *Message) appendFloat64(f float64) {
	m.message = append(m.message, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.BigEndian.PutUint64(m.message[len(m.message)-8:], math.Float64bits(f))
}

func (m *Message) appendUint32(ui uint32) {
	m.message = append(m.message, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(m.message[len(m.message)-4:], ui)
}

func (m *Message) appendBoolean(b bool) {
	if b {
		m.appendSingleByte(1)
	} else {
		m.appendSingleByte(0)
	}
}

func (m *Message) appendBytes(bytes []byte) {
	m.message = append(m.message, bytes...)
}

func (m *Message) appendString(s string) {
	m.appendUint32(uint32(len(s)))
	m.message = append(m.message, []byte(s)...)
}

func (m *Message) appendColor(c color.IColor) {
	m.appendBytes(c.BinaryRGBA())
}

const (
	singleByte int = 1 << iota
	twoBytes
	fourBytes
	eightBytes
)

const (
	emptyString = ""
)

func (m *Message) retrieveSingleByte() byte {
	if len(m.message) < singleByte {
		m.short()
		return 0
	}
	b := m.message[0]
	m.message = m.message[singleByte:]
	return b
}

func (m *Message) retrieveUint32() uint32 {
	if len(m.message) < fourBytes {
		m.short()
		return 0
	}
	i := binary.BigEndian.Uint32(m.message)
	m.message = m.message[fourBytes:]
	return i
}

func (m *Message) retrieveUint64() uint64 {
	if len(m.message) < eightBytes {
		m.short()
		return 0
	}
	i := binary.BigEndian.Uint64(m.message)
	m.message = m.message[eightBytes:]
	return i
}

func (m *Message) retrieveFloat64() float64 {
	return math.Float64frombits(m.retrieveUint64())
}

func (m *Message) retrieveBool() bool {
	return m.retrieveSingleByte() != 0
}

func (m *Message) retrieveString() string {
	length := int(m.retrieveUint32())
	if len(m.message) < length {
		m.short()
		return emptyString
	}
	s := string(m.message[:length])
	m.message = m.message[length:]
	return s
}

func (m *Message) clear() {
	m.message = make([]byte, 0, cap(m.message))
}

func (m *Message) short() {
	m.clear()
	m.err = errDataTooShort{}
}

type errDataTooShort struct{}

func (err errDataTooShort) Error() string {
	return "message bytes array too short to retrieve data"
}
