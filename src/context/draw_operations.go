package context

import (
	"maryprotocol/flags"
	"sakhalin/color"
)

func (ctx *Context) FillStyle(col color.IColor) {
	ctx.AppendBytes(flags.B_FillStyle)
	ctx.AppendColor(col)
}

// gradient

// pattern
