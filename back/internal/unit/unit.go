package unit

import (
	"back/pkg/utils"
	"encoding/json"
	"log"
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
}

func (u *Unit) FillFout(buf []int) {
	len_src := len(buf)
	if len_src > NUMBER_OUTS {
		//log.Printf("len=%d", len_src)
		len_src = NUMBER_OUTS
	}
	for i := 0; i < len_src; i++ {
		u.Fout[i] = buf[i]
	}
}
func (u *Unit) FillFtout(buf []int) {
	len_src := len(buf)
	if len_src > NUMBER_OUTS {
		//log.Printf("len=%d", len_src)
		len_src = NUMBER_OUTS
	}
	for i := 0; i < len_src; i++ {
		u.Ftout[i] = buf[i]
	}
}
func (u *Unit) FillSout(buf []int) {
	len_src := len(buf)
	if len_src > NUMBER_OUTS {
		//.Printf("len=%d", len_src)
		len_src = NUMBER_OUTS
	}
	for i := 0; i < len_src; i++ {
		u.Sout[i] = buf[i]
	}
}
func (u *Unit) FillIndTemper(buf []int) {
	len_src := len(buf)
	if len_src > NUMBER_OUTS {
		//log.Printf("len=%d", len_src)
		len_src = NUMBER_OUTS
	}
	for i := 0; i < len_src; i++ {
		u.IndTemper[i] = buf[i]
	}
}
func (u *Unit) FillTemper(buf []int) {
	len_src := len(buf)
	if len_src > NUMBER_TEMPER {
		//log.Printf("len=%d", len_src)
		len_src = NUMBER_TEMPER
	}
	for i := 0; i < len_src; i++ {
		c := buf[i] & 0xff
		if c != 0x80 && (c&(1<<7)) != 0 {
			c = c - 256
		}
		u.Temper[i] = c
	}
}
func (u *Unit) Fill12v(s string) {
	u.U12v = s
}
func (u *Unit) FillLevelWifi(s int) {
	u.LevelWifi = s
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
