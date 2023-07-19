package mq_mq

import (
	"back/pkg/config"
	"back/pkg/logger"
	"back/pkg/utils"
	"sync/atomic"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	QOS1 = 1
)

type Mq struct {
	Login     string
	password  string
	addr      string
	client    mqtt.Client
	logger    *logger.Logger
	connected atomic.Value
	subOk     atomic.Value
}

func Get(logger *logger.Logger) *Mq {
	cfg := config.Get()
	Login := cfg.MqUser
	addr := cfg.MqHost
	password := cfg.MqPass
	m := Mq{
		Login: Login, addr: addr,
		password:  password,
		client:    nil,
		logger:    logger,
		connected: atomic.Value{},
		subOk:     atomic.Value{},
	}

	m.subOk.Store(false)
	m.connected.Store(false)

	m.InitClient()
	go m.Connect()

	return &m
}

func (m *Mq) InitClient() {
	m.logger.Info().Msg("init mqtt client...")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(m.addr)
	opts.SetClientID(utils.GetClientId())
	opts.SetUsername(m.Login)
	opts.SetPassword(m.password)
	opts.ConnectRetry = false // true
	opts.AutoReconnect = true

	// Log events
	opts.OnConnectionLost = func(cl mqtt.Client, err error) {
		m.logger.Info().Msg("connection lost")
		//m.subOk = false
		m.connected.Store(false)
		m.subOk.Store(false)
	}
	opts.OnConnect = func(c mqtt.Client) {
		m.logger.Info().Msg("connection established")
		m.connected.Store(true)
		//m.Sub(c)
	}
	opts.OnReconnecting = func(mqtt.Client, *mqtt.ClientOptions) {
		m.logger.Info().Msg("attempting to reconnect")
	}
	m.client = mqtt.NewClient(opts)
}

func (m *Mq) Connect() {
	m.logger.Info().Msg("mqtt start connection ... ")
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		//return token.Error()
		//return errors.Wrap(token.Error(), "mqtt connect fail")
		m.logger.Info().Msgf("mqtt connect fail %s", token.Error())
	}
	m.logger.Info().Msg("mqtt ...")
}
func (m *Mq) Disconnect() {
	m.logger.Info().Msg("mqtt disconnect")
	if m.client != nil {
		m.client.Disconnect(1000)
	}
	utils.D_1ms(2)
}

func (m *Mq) Sub(strUnit string, handle func(_ mqtt.Client, msg mqtt.Message)) {
	go func() {
		m.logger.Info().Msgf("sub unit %s, ", strUnit)
		cnt := 0
		if m.client == nil {
			m.logger.Info().Msg(" no client ")
			return
		}
		for { // wait connect to brocker if not connected yet
			if m.connected.Load().(bool) {
				break
			}
			utils.D_1s(1)
			cnt++
			m.logger.Info().Msgf(" mqtt not connected yet, cnt=%d ", cnt)
			if cnt > 9 {
				return
			}
		}
		m.logger.Info().Msgf("start sub %s ", strUnit)

		t := m.client.Subscribe(utils.GetTopicSub(m.Login, strUnit), QOS1, handle)
		//go func() {
		_ = t.Wait() // Can also use '<-t.Done()' in releases > 1.2.0
		if t.Error() != nil {
			m.logger.Info().Msgf("err sub: %s\n", t.Error())
		} else {
			m.logger.Info().Msgf("subscribed to: %s", strUnit)
			m.subOk.Store(true) //m.subOk = true
		}
		//}()
	}()
}

func (m *Mq) IsSubOk() bool {
	//return m.subOk
	subOk := m.subOk.Load().(bool)
	return subOk
}
