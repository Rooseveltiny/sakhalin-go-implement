package context

import "sakhalin/message"

type Context struct {
	settings    Settings
	draws       chan<- []byte
	events      <-chan events.Event
	messagePool message.Message
}

func (ctx *Context) Go() {
	ctx.draws <- ctx.messagePool.Message
}
