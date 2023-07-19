package ws

import (
	"encoding/json"

	"log"
)

type WsSendMes struct {
	Username string `json:"username" binding:"required"` //: this.username,
	Topic    string `json:"topic" binding:"required"`    // "ab@c.ru/5555/devsend/cp",
	Message  string `json:"message" binding:"required"`  // "0001225577",
	Group    string `json:"group" binding:"required"`    // "mqtt"
}

func (h *Hub) SendToWs(user, topic, mes string) {
	wsMes := WsSendMes{
		Username: user,
		Topic:    topic,
		Message:  mes,
		Group:    "mqtt",
	}
	b, err := json.Marshal(wsMes)
	if err != nil {
		log.Printf("err %s", err)
		return
	}

	log.Println(string(b))
	h.broadcast <- b

}

func (h *Hub) SendToWsJson(user, topic, mes string) {
	wsMes := WsSendMes{
		Username: user,
		Topic:    topic,
		Message:  mes,
		Group:    "json",
	}
	b, err := json.Marshal(wsMes)
	if err != nil {
		log.Printf("err %s", err)
		return
	}

	//log.Println(string(b))
	h.broadcast <- b

}
