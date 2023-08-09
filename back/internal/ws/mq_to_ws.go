package ws

import (
	"encoding/json"
	"strings"
	"time"

	"log"
)

type WsSendMes struct {
	Username string `json:"username" binding:"required"` //: this.username,
	Topic    string `json:"topic" binding:"required"`    // "ab@c.ru/5555/devsend/cp",
	Message  string `json:"message" binding:"required"`  // "0001225577",
	Group    string `json:"group" binding:"required"`    // "mqtt"
	Pass     string `json:"pass"`
}

// func (h *Hub) SendToWs(s []string) {
// 	if len(s) < 3 {
// 		log.Println("len < 3")
// 		return
// 	}
// 	wsMes := WsSendMes{
// 		Username: s[0],
// 		Topic:    s[1],
// 		Message:  s[2],
// 		Group:    "mqtt",
// 		Pass:     "-",
// 	}
// 	bytesMes, err := json.Marshal(wsMes)
// 	if err != nil {
// 		log.Printf("err %s", err)
// 		return
// 	}
// 	//log.Println(string(bytesMes))

// 	//h.broadcast <- b
// 	//log.Printf("SendToWs <-client.send ")
// 	for client := range h.clients {
// 		select {
// 		case client.send <- bytesMes:
// 		default:
// 			close(client.send)
// 			delete(h.clients, client)
// 		}
// 	}
// }

func (h *Hub) SendToWsPassErr(s []string) ([]byte, error) {
	//log.Printf("SendToWsJson, s: %v", s)
	wsMes := WsSendMes{
		Username: s[0],
		Topic:    "-",
		Message:  "-",
		Group:    "pass",
		Pass:     "-",
	}
	bytesMes, err := json.Marshal(wsMes)
	if err != nil {
		log.Printf("err %s", err)
		return nil, err
	}
	return bytesMes, nil
}

func (h *Hub) SendToWsJson(s []string) ([]byte, error) {
	//log.Printf("SendToWsJson, s: %v", s)
	if len(s) < 3 {
		log.Println("len < 3")
		return nil, nil
	}
	wsMes := WsSendMes{
		Username: s[0],
		Topic:    s[1],
		Message:  s[2],
		Group:    "json",
		Pass:     "-",
	}
	bytesMes, err := json.Marshal(wsMes)
	if err != nil {
		log.Printf("err %s", err)
		return nil, err
	}
	return bytesMes, nil
}

func (h *Hub) SendToWs(s []string) {
	if len(s) < 3 {
		log.Println("len < 3")
		return
	}
	topic := s[1]
	t := strings.Split(topic, "/")

	strName := t[1]
	s, err := h.service.GetUnitByName(strName, s[0])
	if err != nil {
		log.Printf("err %s", err)
		return
	}
	bytesMes, err := h.SendToWsJson(s)
	if err != nil {
		log.Printf("err %s", err)
		return
	}
	// log.Println("from mq to ws ")
	// log.Println(string(bytesMes))

	for client := range h.clients {
		select {
		case client.send <- bytesMes:
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}
}

func (h *Hub) CheckMes(client *Client, b []byte) {
	wsMes := WsSendMes{}
	err := json.Unmarshal(b, &wsMes)
	if err != nil {
		log.Printf("err %s", err)
		return
	}
	user := wsMes.Username
	group := wsMes.Group
	topic := wsMes.Topic
	message := wsMes.Message
	pass := wsMes.Pass
	if group == "connection" {
		cnt := h.service.GetCountUnits()
		time.Sleep(1 * time.Second) // android don't update
		for i := 0; i < cnt; i++ {
			s, err := h.service.GetUnitByInd(i, user)
			if err != nil {
				log.Printf("err %s", err)
				return
			}
			bytesMes, err := h.SendToWsJson(s)
			if err != nil {
				log.Printf("err %s", err)
				return
			}

			//log.Printf("SendToWsJson <-client.send ")
			select {
			case client.send <- bytesMes:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}
	} else if group == "command" {
		if !h.service.CheckPass(pass) {
			log.Println("wrong pass")
			bytesMes, err := h.SendToWsPassErr([]string{user})
			if err != nil {
				log.Printf("err %s", err)
				return
			}
			select {
			case client.send <- bytesMes:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		} else {
			log.Printf("send to mq %s %s", topic, message)
			s := []string{topic, message}
			h.hglob.WsToMq <- s
		}
	}
}
