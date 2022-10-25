package canvasctx

import (
	"sakhalin/message"
)

type Context struct {
	settings Settings
	Draws    chan []byte
	// events      <-chan events.Event
	messagePool message.Message
}

func NewContext() *Context {
	return &Context{}
}

// func (ctx *Context) Events() <-chan events.Event {
// 	return ctx.events
// }

func (ctx *Context) Go() {
	ctx.Draws <- ctx.messagePool.Message
}
