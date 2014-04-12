package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn

func main() {
	m := martini.Classic()

	m.Use(martini.Static("public"))

	m.Get("/socket", socketHandler)
	m.Post("/new_mail", newMailHandler)

	m.Run()
}

func newMailHandler(r *http.Request, log *log.Logger) int {
	body := r.FormValue("subject")
	log.Println("Mail received: " + body)
	for _, conn := range conns {
		go conn.WriteMessage(1, []byte(body))
	}

	return http.StatusOK
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
	}

	log.Println("WebSocket successfully opened at " + conn.RemoteAddr().String())
	conns = append(conns, conn)
}
