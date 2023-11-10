package mq_mq

import (
	"back/internal/hglob"
	"back/internal/unit"
	"back/pkg/config"
	"back/pkg/logger"
	"back/pkg/utils"
	"context"
	"errors"
	"sync/atomic"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	QOS1 = 1
)

type Mq struct {
	Login     string
	password  string
	addr      string
	us        *unit.Units
	client    mqtt.Client
	logger    *logger.Logger
	hglob     *hglob.Hglob
	connected atomic.Value
	subOk     atomic.Value
	// cmdPub    []ValueWait
	cmdPub map[string]ValueWait
}

func Get(ctx context.Context, logger *logger.Logger, us *unit.Units, hglob *hglob.Hglob) error {
	cfg := config.Get()
	Login := cfg.MqUser
	addr := cfg.MqHost
	password := cfg.MqPass
	if password == "" {
		return errors.New("password is empty")
	}
	m := Mq{
		Login: Login, addr: addr,
		password:  password,
		us:        us,
		client:    nil,
		logger:    logger,
		hglob:     hglob,
		connected: atomic.Value{},
		subOk:     atomic.Value{},
		// cmdPub:    make([]ValueWait, 0),
		cmdPub: make(map[string]ValueWait),
	}

	m.subOk.Store(false)
	m.connected.Store(false)

	m.InitClient()
	go m.Connect()
	go m.WaitToMq(ctx)
	go m.CheckVers(ctx)
	go m.Disconnect(ctx)
	go m.WaitToCmdPub(ctx)
	return nil
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
		m.logger.Info().Msg("mq connection lost")
		m.connected.Store(false)
		m.subOk.Store(false)
	}
	opts.OnConnect = func(c mqtt.Client) {
		m.logger.Info().Msg("mq connected")
		m.connected.Store(true)
		m.SubAll()
	}
	opts.OnReconnecting = func(mqtt.Client, *mqtt.ClientOptions) {
		m.logger.Info().Msg("mq reconnecting")
	}
	m.client = mqtt.NewClient(opts)
}

func (m *Mq) Connect() {
	//m.logger.Info().Msg("mqtt start connection ... ")
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		//return token.Error()
		//return errors.Wrap(token.Error(), "mqtt connect fail")
		m.logger.Info().Msgf("mqtt connect fail %s", token.Error())
	}
	//m.logger.Info().Msg("mqtt ...")
}

func (m *Mq) Disconnect(ctx context.Context) {
	<-ctx.Done()
	//m.logger.Info().Msg("ctx done mq.Disconnect")
	if m.client != nil {
		m.client.Disconnect(1000)
	}
	time.Sleep(2 * time.Millisecond)
	m.logger.Info().Msg("mq disconnected")
}
