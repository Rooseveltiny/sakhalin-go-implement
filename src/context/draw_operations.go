package context

import (
	"sakhalin/color"
	"sakhalin/flags"
)

func (ctx *Context) FillStyle(col color.IColor) {
	ctx.AppendSingleByte(flags.B_FillStyle)
	ctx.AppendColor(col)
}

// gradient

// pattern

func (ctx *Context) Font(font string) {
	ctx.AppendSingleByte(flags.B_Font)
	ctx.AppendString(font)
}

func (ctx *Context) GlobalAlpha(alpha float64) {
	ctx.AppendSingleByte(flags.B_GlobalAlpha)
	ctx.AppendFloat64(alpha)
}

func (ctx *Context) GlobalCompositeOperation(mode flags.CompositeOperation) {
	ctx.AppendSingleByte(flags.B_GlobalCompositeOperation)
	ctx.AppendSingleByte(byte(mode))
}

// image smooth

func (ctx *Context) LineCap(cap flags.LineCap) {
	ctx.AppendSingleByte(flags.B_LineCap)
	ctx.AppendSingleByte(byte(cap))
}

func (ctx *Context) LineDashOffset(offset float64) {
	ctx.AppendSingleByte(flags.B_LineDashOffset)
	ctx.AppendFloat64(offset)
}

func (ctx *Context) LineJoin(join flags.LineJoin) {
	ctx.AppendSingleByte(flags.B_LineJoin)
	ctx.AppendSingleByte(byte(join))
}

func (ctx *Context) LineWidth(width float64) {
	ctx.AppendSingleByte(flags.B_LineWidth)
	ctx.AppendFloat64(width)
}

func (ctx *Context) MiterLimit(value float64) {
	ctx.AppendSingleByte(flags.B_MiterLimit)
	ctx.AppendFloat64(value)
}

func (ctx *Context) ShadowBlur(level float64) {
	ctx.AppendSingleByte(flags.B_ShadowBlur)
	ctx.AppendFloat64(level)
}

func (ctx *Context) ShadowColor(c color.IColor) {
	ctx.AppendSingleByte(flags.B_ShadowColor)
	ctx.AppendColor(c)
}

func (ctx *Context) ShadowOffsetX(offset float64) {
	ctx.AppendSingleByte(flags.B_ShadowOffsetX)
	ctx.AppendFloat64(offset)
}

func (ctx *Context) ShadowOffsetY(offset float64) {
	ctx.AppendSingleByte(flags.B_ShadowOffsetY)
	ctx.AppendFloat64(offset)
}

func (ctx *Context) StrokeStyle(c color.IColor) {
	ctx.AppendSingleByte(flags.B_StrokeStyle)
	ctx.AppendColor(c)
}

// stroke gradient

// stroke pattern

func (ctx *Context) TextAlign(align flags.TextAlign) {
	ctx.AppendSingleByte(flags.B_TextAlign)
	ctx.AppendSingleByte(byte(align))
}

func (ctx *Context) TextBaseLine(baseLine flags.TextBaseline) {
	ctx.AppendSingleByte(flags.B_TextBaseline)
	ctx.AppendSingleByte(byte(baseLine))
}

func (ctx *Context) Arc(x, y, radius, startAngle, endAngle float64, anticlockwise bool) {
	ctx.AppendSingleByte(flags.B_Arc)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
	ctx.AppendFloat64(radius)
	ctx.AppendFloat64(startAngle)
	ctx.AppendFloat64(endAngle)
	ctx.AppendBoolean(anticlockwise)
}

func (ctx *Context) ArcTo(x1, y1, x2, y2, radius float64) {
	ctx.AppendSingleByte(flags.B_ArcTo)
	ctx.AppendFloat64(x1)
	ctx.AppendFloat64(y1)
	ctx.AppendFloat64(x2)
	ctx.AppendFloat64(y2)
	ctx.AppendFloat64(radius)
}

func (ctx *Context) BeginPath() {
	ctx.AppendSingleByte(flags.B_BeginPath)
}

func (ctx *Context) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	ctx.AppendSingleByte(flags.B_BezierCurveTo)
	ctx.AppendFloat64(cp1x)
	ctx.AppendFloat64(cp1y)
	ctx.AppendFloat64(cp2x)
	ctx.AppendFloat64(cp2y)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
}

func (ctx *Context) ClearRect(x, y, width, height float64) {
	ctx.AppendSingleByte(flags.B_ClearRect)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
	ctx.AppendFloat64(width)
	ctx.AppendFloat64(height)
}

func (ctx *Context) Clip() {
	ctx.AppendSingleByte(flags.B_Clip)
}

func (ctx *Context) ClosePath() {
	ctx.AppendSingleByte(flags.B_ClosePath)
}

func (ctx *Context) Ellipse(x, y, radiusX, radiusY, rotation, startAngle, endAngle float64, anticlockwise bool) {
	ctx.AppendSingleByte(flags.B_Ellipse)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
	ctx.AppendFloat64(radiusX)
	ctx.AppendFloat64(radiusY)
	ctx.AppendFloat64(rotation)
	ctx.AppendFloat64(startAngle)
	ctx.AppendFloat64(endAngle)
	ctx.AppendBoolean(anticlockwise)
}

func (ctx *Context) Fill() {
	ctx.AppendSingleByte(flags.B_Fill)
}

func (ctx *Context) FillRect(x, y, width, height float64) {
	ctx.AppendSingleByte(flags.B_FillRect)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
	ctx.AppendFloat64(width)
	ctx.AppendFloat64(height)
}

func (ctx *Context) FillText(text string, x, y float64) {
	ctx.AppendSingleByte(flags.B_FillText)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
	ctx.AppendString(text)
}

func (ctx *Context) FillTextMaxWidth(text string, x, y, maxWidth float64) {
	ctx.AppendSingleByte(flags.B_FillTextMaxWidth)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
	ctx.AppendFloat64(maxWidth)
	ctx.AppendString(text)
}

func (ctx *Context) LineTo(x, y float64) {
	ctx.AppendSingleByte(flags.B_LineTo)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
}

func (ctx *Context) MoveTo(x, y float64) {
	ctx.AppendSingleByte(flags.B_MoveTo)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
}

func (ctx *Context) QuadraticCurveTo(cpx, cpy, x, y float64) {
	ctx.AppendSingleByte(flags.B_QuadraticCurveTo)
	ctx.AppendFloat64(cpx)
	ctx.AppendFloat64(cpy)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
}

func (ctx *Context) Rect(x, y, width, height float64) {
	ctx.AppendSingleByte(flags.B_Rect)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
	ctx.AppendFloat64(width)
	ctx.AppendFloat64(height)
}

func (ctx *Context) Restore() {
	ctx.AppendSingleByte(flags.B_Restore)
}

func (ctx *Context) Rotate(angle float64) {
	ctx.AppendSingleByte(flags.B_Rotate)
	ctx.AppendFloat64(angle)
}

func (ctx *Context) Save() {
	ctx.AppendSingleByte(flags.B_Save)
}

func (ctx *Context) Scale(x, y float64) {
	ctx.AppendSingleByte(flags.B_Scale)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
}

func (ctx *Context) Stroke() {
	ctx.AppendSingleByte(flags.B_Stroke)
}

func (ctx *Context) StrokeText(text string, x, y float64) {
	ctx.AppendSingleByte(flags.B_StrokeText)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
	ctx.AppendString(text)
}

func (ctx *Context) StrokeTextMaxWidth(text string, x, y, maxWidth float64) {
	ctx.AppendSingleByte(flags.B_StrokeTextMaxWidth)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
	ctx.AppendFloat64(maxWidth)
	ctx.AppendString(text)
}

func (ctx *Context) StrokeRect(x, y, width, height float64) {
	ctx.AppendSingleByte(flags.B_StrokeRect)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
	ctx.AppendFloat64(width)
	ctx.AppendFloat64(height)
}

func (ctx *Context) Translate(x, y float64) {
	ctx.AppendSingleByte(flags.B_Translate)
	ctx.AppendFloat64(x)
	ctx.AppendFloat64(y)
}

func (ctx *Context) Transform(a, b, c, d, e, f float64) {
	ctx.AppendSingleByte(flags.B_Transform)
	ctx.AppendFloat64(a)
	ctx.AppendFloat64(b)
	ctx.AppendFloat64(c)
	ctx.AppendFloat64(d)
	ctx.AppendFloat64(e)
	ctx.AppendFloat64(f)
}

func (ctx *Context) PutTransform(a, b, c, d, e, f float64) {
	ctx.AppendSingleByte(flags.B_PutTransform)
	ctx.AppendFloat64(a)
	ctx.AppendFloat64(b)
	ctx.AppendFloat64(c)
	ctx.AppendFloat64(d)
	ctx.AppendFloat64(e)
	ctx.AppendFloat64(f)
}

func (ctx *Context) PutLineDash(segments []float64) {
	ctx.AppendSingleByte(flags.B_PutLineDash)
	ctx.AppendUint32(uint32(len(segments)))
	for _, s := range segments {
		ctx.AppendFloat64(s)
	}
}

// func (ctx *Context) CreateImageData()
// func (ctx *Context) PutImageData()
// func (ctx *Context) PutImageDataDirty()
// func (ctx *Context) DrawImage()
// func (ctx *Context) DrawImageScaled()
// func (ctx *Context) DrawImageSubRectangle()
// func (ctx *Context) CreateLinearGradient()
// func (ctx *Context) CreateRadialGradient()
// func (ctx *Context) CreatePattern()
// func (ctx *Context) GetImageData()
