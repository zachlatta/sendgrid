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

	m.Use(martini.Static("assets"))

	m.Get("/socket", socketHandler)
	m.Post("/new_mail", newMailHandler)

	m.Run()
}

func newMailHandler(w http.ResponseWriter, r *http.Request) int {
	body := r.FormValue("text")
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
	conns = append(conns, conn)
}
