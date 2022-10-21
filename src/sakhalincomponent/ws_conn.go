package sakhalincomponent

import (
	"log"
	"net/http"
	"sakhalin/events"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func readMessages(conn *websocket.Conn, evs chan<- events.Event, wg *sync.WaitGroup) {

}

func writeMessages(conn *websocket.Conn, message <-chan []byte, wg *sync.WaitGroup) {

}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	evs := make(chan events.Event)
	defer close(evs)
	draws := make(chan []byte)
	defer close(draws)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go readMessages(ws, evs, &wg)
	go writeMessages(ws, draws, &wg)

	// ctx := newContext(draws, evs, h.config)
	// go func() {
	// 	defer wg.Done()
	// 	h.draw(ctx)
	// }()

	wg.Wait() // <- blocks performing till wait groups all done
	wg.Add(1)
	evs <- events.StopEvent{}
	wg.Wait()
}

func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}

func RunServe() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
