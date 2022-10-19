package message

import (
	"encoding/binary"
	"math"
	"sakhalin/color"
)

type Message struct {
	Message []byte
	Err     error
}

func NewMessage(b []byte) *Message {
	return &Message{
		Message: b,
		Err:     nil,
	}
}

func (m *Message) AppendSingleByte(b byte) {
	m.Message = append(m.Message, b)
}

func (m *Message) AppendFloat64(f float64) {
	m.Message = append(m.Message, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0)
	binary.BigEndian.PutUint64(m.Message[len(m.Message)-8:], math.Float64bits(f))
}

func (m *Message) AppendUint32(ui uint32) {
	m.Message = append(m.Message, 0x0, 0x0, 0x0, 0x0)
	binary.BigEndian.PutUint32(m.Message[len(m.Message)-4:], ui)
}

func (m *Message) AppendBoolean(b bool) {
	if b {
		m.AppendSingleByte(0x1)
	} else {
		m.AppendSingleByte(0x0)
	}
}

func (m *Message) AppendBytes(bytes []byte) {
	m.Message = append(m.Message, bytes...)
}

func (m *Message) AppendString(s string) {
	m.AppendUint32(uint32(len(s)))
	m.Message = append(m.Message, []byte(s)...)
}

func (m *Message) AppendColor(c color.IColor) {
	m.AppendBytes(c.BinaryRGBA())
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

func (m *Message) RetrieveSingleByte() byte {
	if len(m.Message) < singleByte {
		m.short()
		return 0
	}
	b := m.Message[0]
	m.Message = m.Message[singleByte:]
	return b
}

func (m *Message) RetrieveUint32() uint32 {
	if len(m.Message) < fourBytes {
		m.short()
		return 0
	}
	i := binary.BigEndian.Uint32(m.Message)
	m.Message = m.Message[fourBytes:]
	return i
}

func (m *Message) RetrieveUint64() uint64 {
	if len(m.Message) < eightBytes {
		m.short()
		return 0
	}
	i := binary.BigEndian.Uint64(m.Message)
	m.Message = m.Message[eightBytes:]
	return i
}

func (m *Message) RetrieveFloat64() float64 {
	return math.Float64frombits(m.RetrieveUint64())
}

func (m *Message) RetrieveBool() bool {
	return m.RetrieveSingleByte() != 0
}

func (m *Message) RetrieveString() string {
	length := int(m.RetrieveUint32())
	if len(m.Message) < length {
		m.short()
		return emptyString
	}
	s := string(m.Message[:length])
	m.Message = m.Message[length:]
	return s
}

func (m *Message) Clear() {
	m.Message = make([]byte, 0, cap(m.Message))
}

func (m *Message) short() {
	m.Clear()
	m.Err = errDataTooShort{}
}

type errDataTooShort struct{}

func (err errDataTooShort) Error() string {
	return "Message bytes array too short to retrieve data"
}
