package mq_mq

import (
	//"encoding/json"
	//"back/internal/unit"
	"back/pkg/utils"
	"fmt"
	"log"
	"os"
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

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

type Message struct {
	Count uint64
}

func getClientId() string {
	s := "1111222233334444"
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

// handle is called when a message is received

type Mq struct {
	//us *unit.Units
	login    string
	password string
	addr     string
	client   mqtt.Client
}

//	func (m *Mq) Init(unitPointers []*unit.Unit, cntUnit int) {
//		m.up = make([]*unit.Unit, 10)
//		for i := 0; i < cntUnit; i++ {
//	    m.up[i] = unitPointers[i]
//			//m.up[i].PrintUnit()
//		}
//		m.cnt = cntUnit
//	}
func (m *Mq) Init(addr string, login string, password string) {
	m.login = login
	m.password = password
	m.addr = addr

	m.client = nil
}

// func (m *Mq) recHandle(_ mqtt.Client, msg mqtt.Message) {
// 	topic := msg.Topic()  // ab@m.ru/0803/devsend/
// 	t := strings.Split(topic, "/")
// 	if(t[0] != m.login) {
// 		return
// 	}
// 	mes := string(msg.Payload())
// 	m.us.FillBuf(topic, mes)
// }

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
func (m *Mq) Start() {
	log.Printf("starting mqtt...")
	// Enable logging by uncommenting the below
	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRITICAL] ", 0)
	// mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	// mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	// Create a handler that will deal with incoming messages
	//h := NewHandler()

	// Now we establish the connection to the mqtt broker
	opts := mqtt.NewClientOptions()
	opts.AddBroker(m.addr)
	opts.SetClientID(getClientId())
	opts.SetUsername(m.login)
	opts.SetPassword(m.password)

	//opts.SetOrderMatters(false)       // Allow out of order messages (use this option unless in order delivery is essential)
	//opts.ConnectTimeout = 30*time.Second // Minimal delays on connect
	//opts.WriteTimeout = 30*time.Second   // Minimal delays on writes
	//opts.SetKeepAlive(120*time.Second)              // Keepalive every 10 seconds so we quickly detect network outages
	//opts.PingTimeout = 30*time.Second    // local broker so response should be quick

	// Automate connection management (will keep trying to connect and will reconnect if network drops)
	opts.ConnectRetry = true
	opts.AutoReconnect = true

	// If using QOS2 and CleanSession = FALSE then it is possible that we will receive messages on topics that we
	// have not subscribed to here (if they were previously subscribed to they are part of the session and survive
	// disconnect/reconnect). Adding a DefaultPublishHandler lets us detect this.
	// opts.DefaultPublishHandler = func(_ mqtt.Client, msg mqtt.Message) {
	// 	fmt.Printf("UNEXPECTED MESSAGE: %s\n", msg)
	// }

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

	//
	// Connect to the broker
	//
	m.client = mqtt.NewClient(opts)

	// If using QOS2 and CleanSession = FALSE then messages may be transmitted to us before the subscribe completes.
	// Adding routes prior to connecting is a way of ensuring that these messages are processed
	//client.AddRoute(TOPIC, h.handle)

	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Printf("Connection is up")

	// Messages will be delivered asynchronously so we just need to wait for a signal to shutdown
	// go m.wait_end()
}

func (m *Mq) InitClient() {
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
