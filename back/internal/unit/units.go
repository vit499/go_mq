package unit

import (
	// "back/internal/mq_mq"
	// "back/internal/ws"
	"back/pkg/config"
	"back/pkg/logger"
	"back/pkg/tgbot"
	"context"
	"fmt"

	// "time"

	//"back/pkg/utils"
	"encoding/json"

	//"fmt"
	"log"
	"strings"
	//mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Units struct {
	Up     []*Unit
	Cnt    int
	Tg     *tgbot.Tgbot
	logger *logger.Logger
	pass   string
}

func Get(ctx context.Context, tg *tgbot.Tgbot, logger *logger.Logger) *Units {
	cfg := config.Get()
	arr := cfg.Units
	pass := cfg.MqPass

	up1 := make([]*Unit, 10)
	//us := Units{up1, 0, mq, tg, hub, logger}
	us := Units{up1, 0, tg, logger, pass}
	for i := 0; i < len(arr); i++ {
		us.AddOneUnit(ctx, arr[i])
	}
	//us.Sub()
	//go us.KeepAlive()
	//go us.waitClientWs()
	return &us
}

func (us *Units) AddOneUnit(ctx context.Context, s string) {
	if us.Cnt < 10 {
		ind := us.Cnt
		u := Unit{}
		u.Init(s)
		us.Up[ind] = &u
		us.Cnt = us.Cnt + 1
		us.Up[ind].CheckingOnline(ctx, us.UnitOffline)
	}
}

func (us *Units) CheckPass(p string) bool {
	if us.pass != p {
		us.logger.Info().Msgf("check pass %s mq=%s", p, us.pass)
		return false
	}
	return true
}

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
	restoreOnline := us.Up[indUnit].SetOnline(true)
	us.Up[indUnit].FillBuf(topic, mes)
	mesEvent := us.Up[indUnit].FillBufEv(topic, mes)
	//us.Up[indUnit].PrintUnit()
	if mesEvent != "" {
		us.Tg.SendMes(mesEvent)
	} else {
		if restoreOnline {
			val := fmt.Sprintf("%sR70100000", us.Up[indUnit].StrUnit)
			log.Printf("UnitOnline %s", val)
			mesEvent = us.Up[indUnit].CheckEv(val)
			if mesEvent != "" {
				us.Tg.SendMes(mesEvent)
			}
		}
	}
}

func (us *Units) UnitOffline(s string) {
	indUnit := us.getIndUnit(s)
	if indUnit >= us.Cnt {
		return
	}
	val := fmt.Sprintf("%sE70100000", s)
	log.Printf("UnitOffline %s", val)
	mesEvent := us.Up[indUnit].CheckEv(val)
	if mesEvent != "" {
		us.Tg.SendMes(mesEvent)
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

func (us *Units) GetUnitFtOut(s string) ([]int, bool) {
	indUnit := us.getIndUnit(s)
	if indUnit >= us.Cnt {
		return nil, false
	}
	ftout := us.Up[indUnit].Ftout
	return ftout, true
}
