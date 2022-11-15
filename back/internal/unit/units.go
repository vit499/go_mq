package unit

import (
	"back/internal/mq_mq"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Units struct {
	up   []*Unit
	Cnt  int
	Mq   *mq_mq.Mq
}

func GetUnits(mq *mq_mq.Mq) *Units {
	up1 := make([]*Unit, 10)
	us := Units{up1, 0, mq}
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
		if(us.up[i].StrUnit == s) {
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
	us.up[indUnit].FillBuf(topic, mes)
	us.up[indUnit].PrintUnit()
}
func (us *Units) RecHandle(_ mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()  // ab@m.ru/0803/devsend/
	// t := strings.Split(topic, "/")
	// if(t[0] != m.login) {
	// 	return
	// }
	mes := string(msg.Payload())
	us.FillBuf(topic, mes)
}

func (us *Units) getJsonUnit(s string) string {
	return "1"
}