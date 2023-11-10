package mq_mq

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TypeCmd int64

const (
	CmdSetOut TypeCmd = iota
)

type cmdToMq struct {
	user    string
	nobj    string
	cmd     string
	val     string
	cnt     int
	topic   string
	message string
	enabled bool
}
type ValueWait struct {
	topic   string
	message string
	user    string
	indOut  int
	val     int
	cnt     int
}

// save command to map
func (m *Mq) fillCmdToMq(topic, message string) {
	t := strings.Split(topic, "/") // ab@m.ru/0101/devrec/control
	user := t[0]
	nobj := t[1]
	if !strings.Contains(message, "=") {
		return
	}
	mm := strings.Split(message, "=") // setout1=25
	cmd := mm[0]
	strVal := mm[1]
	if strings.HasPrefix(cmd, "setout") {
		key := fmt.Sprintf("%s/%s", nobj, cmd) // 0802/setout1
		strOut := cmd[6:]
		indOut, err := strconv.Atoi(strOut)
		if err != nil {
			return
		}
		val, err := strconv.Atoi(strVal)
		if err != nil {
			return
		}
		valueWait := ValueWait{
			topic:   topic,
			message: message,
			user:    user,
			indOut:  indOut - 1,
			val:     val,
			cnt:     0,
		}
		m.cmdPub[key] = valueWait
	}
}

func (m *Mq) checkMatchValue(key string, v ValueWait) bool {
	k := strings.Split(key, "/")
	nobj := k[0]
	cmd := k[1]
	if strings.HasPrefix(cmd, "setout") {
		ftout, ok := m.us.GetUnitFtOut(nobj)
		if !ok {
			return true
		}
		res := v.val == ftout[v.indOut]
		m.logger.Info().Msgf("nobj=%s v.out=%d v.val=%d ftout=%v res=%t", nobj, v.indOut, v.val, ftout, res)
		if res {
			return true
		}
	}
	return false
}

// ожидание команды из websocket, запись в map и publish
func (m *Mq) WaitToMq(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//m.logger.Info().Msg("ctx done mq.WaitToMq")
			return
		case msg := <-m.hglob.WsToMq:
			m.logger.Info().Msgf("to mq topic: %s mes: %s\n", msg[0], msg[1])
			m.fillCmdToMq(msg[0], msg[1])
			if m.subOk.Load().(bool) {
				m.client.Publish(msg[0], QOS1, false, msg[1])
			}
		}
	}
}

// если команда еще существует в map то увеличиваем счетчик и publish
func (m *Mq) WaitToCmdPub(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			for key, v := range m.cmdPub {
				if m.checkMatchValue(key, v) {
					delete(m.cmdPub, key)
					continue
				}
				cnt := v.cnt
				cnt++
				if cnt >= 5 {
					delete(m.cmdPub, key)
					continue
				}
				v.cnt = cnt
				m.cmdPub[key] = v
			}
			for _, v := range m.cmdPub {
				m.logger.Info().Msgf("to mq topic: %s mes: %s cnt=%d\n", v.topic, v.message, v.cnt)
				if m.subOk.Load().(bool) {
					m.client.Publish(v.topic, QOS1, false, v.message)
				}
			}
		}
	}
}
