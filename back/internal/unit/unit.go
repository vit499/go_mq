package unit

import (
	"back/pkg/utils"
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"
)

const NUMBER_OUTS = 4
const NUMBER_TEMPER = 3

type Unit struct {
	StrUnit   string `json:"StrUnit" binding:"required"`
	Fout      []int  `json:"Fout" binding:"required"`      // function of outs
	Sout      []int  `json:"Sout" binding:"required"`      // status outs
	Ftout     []int  `json:"Ftout" binding:"required"`     // temper for ON-OFF
	IndTemper []int  `json:"IndTemper" binding:"required"` // index of tempers for out
	Temper    []int  `json:"Temper" binding:"required"`    // value tempers
	U12v      string `json:"U12v" binding:"required"`
	LevelGsm  int    `json:"LevelGsm" binding:"required"`
	LevelWifi int    `json:"LevelWifi" binding:"required"`
	Vers      string `json:"Vers"`
	Online    bool   `json:"Online"`
	cnt       int
	mutex     sync.Mutex
}

func (u *Unit) Init(strUnit string) {
	u.StrUnit = strUnit
	u.Fout = make([]int, NUMBER_OUTS)
	u.Sout = make([]int, NUMBER_OUTS)
	u.Ftout = make([]int, NUMBER_OUTS)
	u.IndTemper = make([]int, NUMBER_OUTS)
	u.Temper = make([]int, NUMBER_TEMPER)
	for i := 0; i < NUMBER_OUTS; i++ {
		u.Fout[i] = 0
		u.Sout[i] = 0
		u.Ftout[i] = 0
		u.IndTemper[i] = 0
	}
	for i := 0; i < NUMBER_TEMPER; i++ {
		u.Temper[i] = 0x80
	}
	u.U12v = "-"
	u.LevelGsm = 0
	u.LevelWifi = 0
	u.Vers = ""
	u.Online = false
	u.cnt = 0
	u.mutex = sync.Mutex{}
	//go u.checkOnline()
}

func (u *Unit) CheckingOnline(ctx context.Context) {
	go u.checkOnline(ctx)
	// ticker := time.NewTicker(1 * time.Second)
	// defer ticker.Stop()
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		return
	// 	case <-ticker.C:
	// 		u.checkOnline()
	// 	}
	// }
}

func (u *Unit) PrintUnit() {
	s := u.StrUnit
	log.Printf("Unit= %s", s)
	s = utils.Hex2Str(u.Fout, NUMBER_OUTS)
	log.Printf("Fout= %s", s)
	s = utils.Hex2Str(u.Sout, NUMBER_OUTS)
	log.Printf("Sout= %s", s)
	s = utils.Hex2Str(u.Ftout, NUMBER_OUTS)
	log.Printf("Ftout= %s", s)
	s = utils.Hex2Str(u.IndTemper, NUMBER_OUTS)
	log.Printf("IndTemper= %s", s)
	s = utils.Hex2Str(u.Temper, NUMBER_TEMPER)
	log.Printf("Temper= %s", s)
	log.Printf("U12v=%s", u.U12v)
	log.Printf("WiFi=%d", u.LevelWifi)

	b, err := json.Marshal(u)
	if err != nil {
		log.Printf("err %s", err)
	}
	log.Printf("b=%s", b)
	if u.Temper[0] != 0x80 {
		log.Printf("T1=%d", u.Temper[0])
	}
	if u.Temper[1] != 0x80 {
		log.Printf("T2=%d", u.Temper[1])
	}
}

func (u *Unit) SetOnline(v bool) {
	//log.Printf("online %s %t", u.StrUnit, v)
	u.mutex.Lock()
	u.Online = v
	u.cnt = 0
	u.mutex.Unlock()
}
func (u *Unit) checkOnline(ctx context.Context) {
	for {
		ticker := time.NewTicker(60 * time.Second)
		select {
		case <-ctx.Done():
			log.Printf("ctx done checkOnline")
			return
		case <-ticker.C:
			u.mutex.Lock()
			if u.Online {
				u.cnt++
				if u.cnt >= 10 {
					u.Online = false
				}
			}
			u.mutex.Unlock()
		}
		// time.Sleep(60 * time.Second)
		// u.mutex.Lock()
		// if u.Online {
		// 	u.cnt++
		// 	if u.cnt >= 10 {
		// 		u.Online = false
		// 	}
		// }
		// u.mutex.Unlock()
	}
}
