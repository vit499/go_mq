package ws

import (
	"back/internal/hglob"
	"back/internal/service/units_service"
	"context"
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

func NewHub(ctx context.Context, service *service.UnitsService, hglob *hglob.Hglob) *Hub {
	h := Hub{
		//broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		//Reg:        make(chan *Client),
		hglob:   hglob,
		service: service,
	}
	go h.Run(ctx)
	return &h
}

func (h *Hub) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// log.Println("ctx done hub.Run")
			return
		case client := <-h.register:
			//log.Println("rec from <-register")
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case msg := <-h.hglob.MqToWs:
			h.SendToWs(msg)
		}
	}
}
