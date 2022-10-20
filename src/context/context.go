package context

import (
	"sakhalin/events"
	"sakhalin/message"
)

type Context struct {
	settings    Settings
	draws       chan<- []byte
	events      <-chan events.Event
	messagePool message.Message
}

func (ctx *Context) Events() <-chan events.Event {
	return ctx.events
}

func (ctx *Context) Go() {
	ctx.draws <- ctx.messagePool.Message
}
