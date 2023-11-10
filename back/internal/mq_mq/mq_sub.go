package mq_mq

import (
	"back/pkg/utils"
	"context"
	"fmt"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

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

func (m *Mq) CheckVers(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
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
