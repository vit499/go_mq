package mq_mq

import (
	//"encoding/json"
	//"back/internal/unit"
	"back/pkg/utils"
	"fmt"
	"log"
	"math/rand"

	//"os"
	"time"

	// "os/signal"
	// "syscall"
	//"strings"
	//"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	QOS1 = 1
	//MQTT_ADDR   = "tcp://vit496.ru:2084"
	WRITETOLOG  = true  // If true then received messages will be written to the console
	WRITETODISK = false // If true then received messages will be written to the file below
	OUTPUTFILE  = "/binds/receivedMessages.txt"
)

// type handler struct {
// }

// func NewHandler() *handler {
// 	return &handler{}
// }

// type Message struct {
// 	Count uint64
// }

func getClientId() string {
	rand.Seed(time.Now().UnixNano())
	r1 := rand.Intn(65000)
	r2 := rand.Intn(65000)
	fmt.Printf("\r\n r2=%X%X", r1, r2)
	s := fmt.Sprintf("%s%X%X", "22223333", r1, r2)
	return s
}
func (m *Mq) getTopicSub(unit string) string {
	s := fmt.Sprintf("%s/%s/devsend/#", m.login, unit)
	//log.Printf("sub topic: %s ", s)
	return s
}

// func getTopicPub(unit string) string {
// 	s := fmt.Sprintf("%s/%s/devrec/control", LOGIN, unit)
// 	return s
// }

type Mq struct {
	login    string
	password string
	addr     string
	client   mqtt.Client
}

func (m *Mq) InitClient(addr string, login string, password string) {
	m.login = login
	m.password = password
	m.addr = addr
	m.client = nil
	log.Printf("init mqtt client...")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(m.addr)
	opts.SetClientID(getClientId())
	opts.SetUsername(m.login)
	opts.SetPassword(m.password)
	opts.ConnectRetry = false // true
	opts.AutoReconnect = true

	// Log events
	opts.OnConnectionLost = func(cl mqtt.Client, err error) {
		log.Printf("connection lost")
	}
	opts.OnConnect = func(c mqtt.Client) {
		log.Printf("connection established")
		//m.Sub(c)
	}
	opts.OnReconnecting = func(mqtt.Client, *mqtt.ClientOptions) {
		log.Printf("attempting to reconnect")
	}
	m.client = mqtt.NewClient(opts)
}

func (m *Mq) Connect() error {
	log.Printf("mqtt start connection ... ")
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		//panic(token.Error())
		return token.Error()
	}
	log.Printf("mqtt ...")
	return nil
}
func (m *Mq) Disconnect() {
	log.Printf("mqtt disconnect")
	if m.client != nil {
		m.client.Disconnect(1000)
	}
	utils.D_1ms(2)
}

func (m *Mq) Sub(strUnit string, handle func(_ mqtt.Client, msg mqtt.Message)) {
	go func() {
		log.Printf("sub unit %s, ", strUnit)
		cnt := 0
		if m.client == nil {
			log.Printf(" no client ")
			return
		}
		for {
			if m.client.IsConnected() {
				break
			}
			utils.D_1s(1)
			cnt++
			log.Printf(" mqtt not connected %d ", cnt)
			if cnt > 9 {
				return
			}
		}
		log.Printf("start sub %s ", strUnit)

		t := m.client.Subscribe(m.getTopicSub(strUnit), QOS1, handle)
		//go func() {
		_ = t.Wait() // Can also use '<-t.Done()' in releases > 1.2.0
		if t.Error() != nil {
			log.Printf("err sub: %s\n", t.Error())
		} else {
			log.Printf("subscribed to: %s", strUnit)
		}
		//}()
	}()
}

// func (m *Mq) wait_end () {
// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig, os.Interrupt)
// 	signal.Notify(sig, syscall.SIGTERM)

// 	<-sig
// 	log.Printf("signal caught - exiting")
// 	m.client.Disconnect(1000)
// 	log.Printf("shutdown complete")
// }
