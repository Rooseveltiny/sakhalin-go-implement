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
