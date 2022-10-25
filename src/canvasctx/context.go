package canvasctx

import (
	"sakhalin/message"
)

type Context struct {
	settings    Settings
	Draws       chan []byte
	messagePool message.Message
}

func NewContext() *Context {
	return &Context{
		settings:    Settings{},
		Draws:       make(chan []byte),
		messagePool: message.Message{},
	}
}

func (ctx *Context) Draw() {
	ctx.Draws <- ctx.messagePool.Message
	ctx.Clear()
}
