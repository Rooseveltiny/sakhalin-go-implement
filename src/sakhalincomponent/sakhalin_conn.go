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
}

func (sC *SakhalinConnection) RetrieveContext() *canvasctx.Context {
	return sC.Context
}

func (sC *SakhalinConnection) RetrieveEvents(eventCh chan events.Event) {
	go func() {
		eventData := <-sC.WSConn.readCh
		decodedEvent, _ := events.DecodeEvent(message.NewMessage(eventData))
		eventCh <- decodedEvent
	}()
}

func (sC *SakhalinConnection) Draw(ctx context.Context) {
	go func() {
		for {
			select {
			case draws := <-sC.Context.Draws:
				sC.WSConn.writeCh <- draws
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (sC *SakhalinConnection) RunSakhalinListen(ctx context.Context) {
	sC.Draw(ctx)
	sC.WSConn.RunServe(ctx)
}

func NewSakhalinConnection(listenRoute string) *SakhalinConnection {
	newWS := NewWSConn("/ws")
	canvasCtx := canvasctx.NewContext()
	newSakhalin := &SakhalinConnection{
		Context: canvasCtx,
		WSConn:  newWS,
		// Events:  make(chan events.Event),
	}
	return newSakhalin
}
