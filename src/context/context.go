package context

import "sakhalin/message"

type Context struct {
	settings Settings
	draws    chan<- []byte
	// events <- chan
	messagePool message.Message
}
