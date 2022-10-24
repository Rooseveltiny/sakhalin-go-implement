package sakhalincomponent

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	. "github.com/smartystreets/goconvey/convey"
)

func ClientWSConn(ctx context.Context) {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

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

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t)
			err := c.WriteMessage(websocket.BinaryMessage, []byte("Hola!"))
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

func TestWsConn(t *testing.T) {
	Convey("test init of ws conn", t, func() {
		wsconn := NewWSConn()
		So(wsconn.wsconn, ShouldBeNil)
		Convey("run, serve and get Hola message", func(c C) {
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			go func() {
				time.Sleep(1000 * time.Millisecond)
				ClientWSConn(ctx)
			}()
			go func() {
				for {
					message := <-wsconn.readCh
					fmt.Printf("message succesfully gotten: %s", string(message))
					c.So(string(message), ShouldEqual, "Hola!")
					cancel()
				}
			}()
			wsconn.RunServe(ctx)
		})
	})
}
