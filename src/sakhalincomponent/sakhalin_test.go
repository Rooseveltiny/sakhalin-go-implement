package sakhalincomponent

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"sakhalin/events"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	. "github.com/smartystreets/goconvey/convey"
)

func SendMessageViaWebSocket(ctx context.Context, sendMessage chan []byte) {

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8081", Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	fmt.Println("succesfully conected")
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	for {
		select {
		case m := <-sendMessage:
			fmt.Println("starting to send message")
			err := c.WriteMessage(websocket.BinaryMessage, m)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-ctx.Done():
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			return
		}
	}
}

func TestSakhalinSendingMessages(t *testing.T) {
	Convey("init sakhalin", t, func(c C) {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		sakhalinConn := NewSakhalinConnection("/ws")
		Convey("test sakhalin send event", func(c C) {
			go func() {
				// init sakhalin connection
				evs := make(chan events.Event)
				sakhalinConn.RetrieveEvents(evs)
				fmt.Println(<-evs)
			}()
			go func() {
				// init client connection
				time.Sleep(1 * time.Second)
				sendM := make(chan []byte)
				go SendMessageViaWebSocket(ctx, sendM)
				time.Sleep(1 * time.Second)
				sendM <- []byte{9, 0b00000001, 0x0, 0x0, 0x11, 0x95, 0x0, 0x0, 0x2, 0xBC, 0b0001000}
			}()
			Convey("test send draw operations", func(c C) {
				go func() {
					canvasCtx := sakhalinConn.RetrieveContext()
					canvasCtx.LineTo(10, 10)
					canvasCtx.Go()
					fmt.Println(canvasCtx)
				}()
				go func() {
					// init closing gorutines
					time.Sleep(5 * time.Second)
					cancel()
				}()
				sakhalinConn.RunSakhalinListen(ctx)
			})
		})
	})

}

func TestSakhalinInit(t *testing.T) {
	Convey("init test sakhalin and gracefull shutdown", t, func(c C) {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		go func() {
			time.Sleep(50 * time.Millisecond)
			cancel()
		}()
		sc := NewSakhalinConnection("/ws")
		sc.RunSakhalinListen(ctx)
	})
}
