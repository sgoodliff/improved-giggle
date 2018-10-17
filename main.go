package main

import (
	"bytes"
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"

	"github.com/sgoodliff/improved-giggle/balance"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}

	addr     = flag.String("addr", "localhost:8080", "http service address")
	upgrader = websocket.Upgrader{} // use default options
)

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	b := []byte("Hello, goodbye, etc!")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {

		b = bytes.TrimSpace(bytes.Replace(b, newline, space, -1))
		log.Printf("recv: %s", b)
		err = c.WriteMessage(websocket.TextMessage, b)
		if err != nil {
			log.Println("write:", err)
			break
		}
		time.Sleep(1000 * time.Millisecond)

	}
}
func updateData() {
	balances := make(map[int]int)
	var mybalance int
	for {
		for i := 0; i < 100; i++ {
			mybalance = balance.GetBalance(i)
			if mybalance != balances[i] {
				log.Debug("Current balance has changed")
				balances[i] = mybalance
				// we need to push this to client
				balance.PushBalance(i, mybalance)
			}
		}
		time.Sleep(10000 * time.Millisecond)
	}
}

func main() {
	hub := newHub()
	go hub.run()
	log.Info("Hello, World")
	//go updateData()
	http.HandleFunc("/home", balanceHomeHandler)
	//	http.HandleFunc("/ws", balanceHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
