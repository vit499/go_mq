package hglob

//import "back/internal/ws"

type Hglob struct {
	MqToWs chan []string
	//HtToWs chan []string
	WsToMq chan []string
}

func NewHglob() *Hglob {
	h := Hglob{
		MqToWs: make(chan []string),
		//HtToWs: make(chan []string),
		WsToMq: make(chan []string),
	}
	return &h
}
