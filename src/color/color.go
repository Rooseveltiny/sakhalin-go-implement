package color

type IColor interface {
	BinaryRGBA() []byte
}

type Color struct {
	R, G, B, A uint8
}

func (c Color) BinaryRGBA() []byte {
	return []byte{c.R, c.G, c.B, c.A}
}

func NewColor(r, g, b, a uint8) IColor {
	return &Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
