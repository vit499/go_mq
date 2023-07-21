package ws

import (
	"back/internal/hglob"
	"back/internal/service/units_service"
	//"log"
)

// "bytes"
// "log"
// "net/http"
// "time"

// "github.com/gorilla/websocket"

type Hub struct {
	clients map[*Client]bool
	//broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	//Reg        chan *Client
	hglob   *hglob.Hglob
	service *service.UnitsService
}

func NewHub(service *service.UnitsService, hglob *hglob.Hglob) *Hub {
	h := Hub{
		//broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		//Reg:        make(chan *Client),
		hglob:   hglob,
		service: service,
	}
	go h.Run()
	return &h
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			//log.Println("rec from <-register")
			h.clients[client] = true
		case client := <-h.unregister:
			//log.Println("rec from <-unregister")
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		// case msg := <-h.hglob.HtToWs:
		// 	log.Println("rec from <-HtToWs ")
		// 	h.SendToWsJson(msg)
		case msg := <-h.hglob.MqToWs:
			//log.Println("rec from <-MqToWs ")
			h.SendToWs(msg)
			// case message := <-h.broadcast:
			// 	//log.Printf("send msg: %s", string(message))
			// 	for client := range h.clients {
			// 		select {
			// 		case client.send <- message:
			// 		default:
			// 			close(client.send)
			// 			delete(h.clients, client)
			// 		}
			// 	}
		}
	}
}
