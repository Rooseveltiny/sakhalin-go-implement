package sakhalincomponent

import "sakhalin/context"

/*
	Sakhalin 'WebSocket' connection is a heart component of Sakhalin canvas protocol.
	It can be initialized and use to communicate with canvas context on a client side.
*/

type SakhalinConnection struct {
	Context context.Context
}

func (sC *SakhalinConnection) RetrieveContext() *context.Context {
	return &sC.Context
}
