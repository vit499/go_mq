package unit

import (
	// "back/internal/mq_mq"
	// "back/internal/ws"
	"back/pkg/config"
	"back/pkg/logger"
	"back/pkg/tgbot"
	// "time"

	//"back/pkg/utils"
	"encoding/json"

	//"fmt"
	"log"
	"strings"
	//mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Units struct {
	Up  []*Unit
	Cnt int
	//mq     *mq_mq.Mq
	Tg *tgbot.Tgbot
	//hub    *ws.Hub
	logger *logger.Logger
}

func Get(tg *tgbot.Tgbot, logger *logger.Logger) *Units {
	cfg := config.Get()
	arr := cfg.Units

	up1 := make([]*Unit, 10)
	//us := Units{up1, 0, mq, tg, hub, logger}
	us := Units{up1, 0, tg, logger}
	for i := 0; i < len(arr); i++ {
		us.AddOneUnit(arr[i])
	}
	//us.Sub()
	//go us.KeepAlive()
	//go us.waitClientWs()
	return &us
}

//	func GetUnits(mq *mq_mq.Mq, tg *tgbot.Tgbot) *Units {
//		up1 := make([]*Unit, 10)
//		us := Units{up1, 0, mq, tg}
//		return &us
//	}
func (us *Units) AddOneUnit(s string) {
	if us.Cnt < 10 {
		ind := us.Cnt
		u := Unit{}
		u.Init(s)
		us.Up[ind] = &u
		us.Cnt = us.Cnt + 1
	}
}

// func (us *Units) Sub() {
// 	for i := 0; i < us.Cnt; i++ {
// 		us.mq.Sub(us.Up[i].StrUnit, us.RecHandle)
// 	}
// }

// func (us *Units) KeepAlive() {

// 	for {
// 		time.Sleep(20 * time.Second)
// 		if !us.mq.IsSubOk() {
// 			us.logger.Info().Msg("sub lost, reconnect...")
// 			us.Sub()
// 		}
// 	}
// }

// func (us *Units) waitClientWs() {
// 	for {
// 		select {
// 		case client := <-us.hub.Reg:
// 			us.logger.Info().Msgf("new client, send ws json %v", client)
// 			us.FormToWsJson()
// 		}
// 	}
// }

func (us *Units) getIndUnit(s string) int {
	r := 100
	for i := 0; i < us.Cnt; i++ {
		if us.Up[i].StrUnit == s {
			return i
		}
	}
	return r
}
func (us *Units) FillBuf(topic string, mes string) {
	//topic := msg.Topic()  // ab@m.ru/0803/devsend/
	//log.Printf("topic= %s, msg= %s ", topic, mes)
	t := strings.Split(topic, "/")
	indUnit := us.getIndUnit(t[1])
	if indUnit >= us.Cnt {
		return
	}
	topic = strings.Join(t[3:], "/")
	//log.Printf("topic= %s, msg= %s ", topic, mes)
	us.Up[indUnit].FillBuf(topic, mes)
	mesEvent := us.Up[indUnit].FillBufEv(topic, mes)
	//us.Up[indUnit].PrintUnit()
	if mesEvent != "" {
		//us.Tg.SendMes(mesEvent)
	}
}

func (us *Units) RecHandle(topic, mes string) {
	us.FillBuf(topic, mes)
	//SendWs(us.hub, user, topic, mes)
}

// func (us *Units) FormToWsJson() {
// 	for ind := 0; ind < us.Cnt; ind++ {
// 		user := us.mq.Login
// 		topic := us.Up[ind].StrUnit
// 		b, err := us.GetJsonUnit(ind)
// 		if err != nil {
// 			//
// 			continue
// 		}
// 		SendWsJson(us.hub, user, topic, string(b))
// 	}
// }

func (us *Units) GetUnitTemper(ind int) ([]int, error) {
	u := us.Up[ind]
	return u.Temper, nil
}

func (us *Units) GetJsonUnit(ind int) ([]byte, error) {
	u := us.Up[ind]
	b, err := json.Marshal(u)
	if err != nil {
		log.Printf("err %s", err)
		return nil, err
	}
	return b, nil
}
func (us *Units) getJsonUnits(s string) string {
	return "1"
}

func (us *Units) GetUnitVers(ind int) string {
	return us.Up[ind].Vers
}
