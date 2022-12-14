package unit

import (
	"back/internal/mq_mq"
	"back/pkg/tgbot"
	//"back/pkg/utils"
	"encoding/json"

	//"fmt"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Units struct {
	up  []*Unit
	Cnt int
	Mq  *mq_mq.Mq
	Tg  *tgbot.Tgbot
}

func GetUnits(mq *mq_mq.Mq, tg *tgbot.Tgbot) *Units {
	up1 := make([]*Unit, 10)
	us := Units{up1, 0, mq, tg}
	return &us
}
func (us *Units) AddUnit(s string) {
	if us.Cnt < 10 {
		ind := us.Cnt
		u := Unit{}
		u.Init(s)
		us.up[ind] = &u
		us.Cnt = us.Cnt + 1
		log.Printf("add unit %s, cnt=%d ", s, us.Cnt)
		us.Mq.Sub(us.up[ind].StrUnit, us.RecHandle)
	}
}

func (us *Units) getIndUnit(s string) int {
	r := 100
	for i := 0; i < us.Cnt; i++ {
		if us.up[i].StrUnit == s {
			return i
		}
	}
	return r
}
func (us *Units) FillBuf(topic string, mes string) {
	//topic := msg.Topic()  // ab@m.ru/0803/devsend/
	t := strings.Split(topic, "/")
	indUnit := us.getIndUnit(t[1])
	if indUnit >= us.Cnt {
		return
	}
	topic = strings.Join(t[3:], "/")
	log.Printf("topic= %s, msg= %s ", topic, mes)
	mesEvent := us.up[indUnit].FillBuf(topic, mes)
	//us.up[indUnit].PrintUnit()
	if mesEvent != "" {
		us.Tg.SendMes(mesEvent)
	}
}
func (us *Units) RecHandle(_ mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic() // ab@m.ru/0803/devsend/
	// t := strings.Split(topic, "/")
	// if(t[0] != m.login) {
	// 	return
	// }
	mes := string(msg.Payload())
	us.FillBuf(topic, mes)
}

func (us *Units) GetUnitTemper(ind int) ([]int, error) {
	var u *Unit
	u = us.up[ind]
	b := u.Temper
	return b, nil
}
func (us *Units) GetJsonUnit(ind int) ([]byte, error) {
	var u *Unit
	u = us.up[ind]
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
