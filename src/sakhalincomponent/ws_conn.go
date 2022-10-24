package sakhalincomponent

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WSConn struct {
	route   string
	wsconn  *websocket.Conn
	readCh  chan []byte
	writeCh chan []byte
}

func (wsc *WSConn) WSPerformer(w http.ResponseWriter, r *http.Request) {
	var err error
	wsc.wsconn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer close(wsc.writeCh)
	defer close(wsc.readCh)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	go wsc.writeMessage(&waitGroup)
	go wsc.readMessage(&waitGroup)
	waitGroup.Wait()
	// waitGroup.Add(1)
	// wsc.readCh <- []byte{0b00000000}
	// waitGroup.Wait()
}

func (wsc *WSConn) readMessage(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		messageType, p, err := wsc.wsconn.ReadMessage()
		if err != nil {
			break
		}
		if messageType != websocket.BinaryMessage {
			continue
		}
		wsc.readCh <- p
	}
}

func (wsc *WSConn) writeMessage(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		message := <-wsc.writeCh
		err := wsc.wsconn.WriteMessage(websocket.BinaryMessage, message)
		if err != nil {
			break
		}
	}
}

func (wsc *WSConn) RunServe(ctx context.Context) {
	s := &http.Server{
		Addr: ":8081",
	}
	http.HandleFunc(wsc.route, wsc.WSPerformer)
	go func() {
		for {
			<-ctx.Done()
			s.Shutdown(ctx)
		}
	}()
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func NewWSConn() *WSConn {
	return &WSConn{
		route:   "/ws",
		readCh:  make(chan []byte),
		writeCh: make(chan []byte),
	}
}
