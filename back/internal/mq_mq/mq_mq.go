package mq_mq

import (
	"back/internal/hglob"
	"back/internal/unit"
	"back/pkg/config"
	"back/pkg/logger"
	"back/pkg/utils"
	"context"
	"errors"
	"fmt"
	"strings"
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
	}

	m.subOk.Store(false)
	m.connected.Store(false)

	m.InitClient()
	go m.Connect()
	go m.WaitToMq(ctx)
	go m.CheckVers(ctx)
	go m.Disconnect(ctx)
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

func (m *Mq) Sub(strUnit string) {
	go func() {
		//m.logger.Info().Msgf("sub unit %s ", strUnit)
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
		//m.logger.Info().Msgf("start sub %s ", strUnit)

		//t := m.client.Subscribe(utils.GetTopicSub(m.Login, strUnit), QOS1, handle)
		t := m.client.Subscribe(utils.GetTopicSub(m.Login, strUnit), QOS1, m.rr)
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

type RecHandler interface {
	RecHandle(topic, mes string)
}

func (m *Mq) recHandle(topic, mes string) {
	m.us.RecHandle(topic, mes)
	t := strings.Split(topic, "/")
	user := t[0]
	s := []string{user, topic, mes}
	m.hglob.MqToWs <- s
}

func (m *Mq) rr(_ mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	message := string(msg.Payload())
	m.recHandle(topic, message)
}

func (m *Mq) SubAll() {
	//m.logger.Info().Msg("sub all")
	for i := 0; i < m.us.Cnt; i++ {
		m.Sub(m.us.Up[i].StrUnit)
	}
}

func (m *Mq) IsSubOk() bool {
	//return m.subOk
	subOk := m.subOk.Load().(bool)
	return subOk
}

func (m *Mq) WaitToMq(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//m.logger.Info().Msg("ctx done mq.WaitToMq")
			return
		case msg := <-m.hglob.WsToMq:
			m.logger.Info().Msgf("to mq topic: %s mes: %s\n", msg[0], msg[1])
			if m.subOk.Load().(bool) {
				m.client.Publish(msg[0], QOS1, false, msg[1])
			}
			// _ = t.Wait() // Can also use '<-t.Done()' in releases > 1.2.0
			// if t.Error() != nil {
			// 	m.logger.Info().Msgf("err pub: %s\n", t.Error())
			// } else {
			// 	m.logger.Info().Msgf("publish ok")
			// }
		}
	}
}

func (m *Mq) CheckVers(ctx context.Context) {

	for {
		//time.Sleep(30 * time.Second)
		ticker := time.NewTicker(30 * time.Second)
		select {
		case <-ctx.Done():
			//m.logger.Info().Msg("ctx done mq.CheckVers")
			return
		case <-ticker.C:
			if m.subOk.Load().(bool) {
				cnt := m.us.Cnt
				for i := 0; i < cnt; i++ {
					vers := m.us.GetUnitVers(i)
					if vers == "" {
						topic := fmt.Sprintf("%s/%s/devrec/control", m.Login, m.us.Up[i].StrUnit)
						message := "reqconfig"
						//m.logger.Info().Msgf("to mq topic: %s mes: %s\n", topic, message)
						m.client.Publish(topic, QOS1, false, message)
					}
				}
			}
		}

	}
}
