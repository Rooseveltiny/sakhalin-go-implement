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
