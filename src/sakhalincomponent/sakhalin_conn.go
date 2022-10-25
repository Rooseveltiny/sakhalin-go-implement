package sakhalincomponent

import (
	"context"
	"sakhalin/canvasctx"
	"sakhalin/events"
	"sakhalin/message"
)

/*
	Sakhalin 'WebSocket' connection is a heart component of Sakhalin canvas protocol.
	It can be initialized and use to communicate with canvas context on a client side.
*/

type SakhalinConnection struct {
	Context *canvasctx.Context
	WSConn  *WSConn
	Events  chan events.Event
}

func (sC *SakhalinConnection) RetrieveContext() *canvasctx.Context {
	return sC.Context
}

func (sC *SakhalinConnection) RetrieveEvents(eventCh chan events.Event) {
	eventData := <-sC.WSConn.readCh
	decodedEvent, _ := events.DecodeEvent(message.NewMessage(eventData))
	eventCh <- decodedEvent
}

func (sC *SakhalinConnection) Draw() {
	sC.WSConn.writeCh <- <-sC.RetrieveContext().Draws
}

func NewSakhalinConnection(listenRoute string, ctx context.Context) *SakhalinConnection {
	newWS := NewWSConn("/ws")
	canvasCtx := canvasctx.NewContext()
	newSakhalin := &SakhalinConnection{
		Context: canvasCtx,
		WSConn:  newWS,
		Events:  make(chan events.Event),
	}
	newSakhalin.WSConn.RunServe(ctx)
	return newSakhalin
}
