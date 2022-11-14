package unit

import (
	"back/pkg/utils"
	"encoding/json"
	"log"
)

const NUMBER_OUTS = 32
const NUMBER_TEMPER = 32

type Unit struct {
	StrUnit string
	Fout []uint8       // function of outs
	Sout []uint8       // status outs
	Ftout []uint8      // temper for ON-OFF
	IndTemper []uint8  // index of tempers for out
	Temper []uint8     // value tempers
	U12v string
	LevelGsm int
	LevelWifi int
}

func (u *Unit) Init(strUnit string) {
	u.StrUnit = strUnit
	u.Fout = make([]uint8, NUMBER_OUTS)
	u.Sout = make([]uint8, NUMBER_OUTS)
	u.Ftout = make([]uint8, NUMBER_OUTS)
	u.IndTemper = make([]uint8, NUMBER_OUTS)
	u.Temper = make([]uint8, NUMBER_TEMPER)
	for i := 0; i < NUMBER_OUTS; i++ { 
		u.Fout[i] = 0
		u.Sout[i] = 0
		u.Temper[i] = 0x80
	}
	u.U12v = "-"
	u.LevelGsm = 0
	u.LevelWifi = 0
}

func (u *Unit) FillFout(buf []uint8) {
	len_src := len(buf)
	if(len_src > NUMBER_OUTS) {
		log.Printf("len=%d", len_src)
		len_src = NUMBER_OUTS
	}
	for i := 0; i < len_src; i++ {
		u.Fout[i] = buf[i]
	}
}
func (u *Unit) FillFtout(buf []uint8) {
	len_src := len(buf)
	if(len_src > NUMBER_OUTS) {
		log.Printf("len=%d", len_src)
		len_src = NUMBER_OUTS
	}
	for i := 0; i < len_src; i++ {
		u.Ftout[i] = buf[i]
	}
}
func (u *Unit) FillSout(buf []uint8) {
	len_src := len(buf)
	if(len_src > NUMBER_OUTS) {
		log.Printf("len=%d", len_src)
		len_src = NUMBER_OUTS
	}
	for i := 0; i < len_src; i++ {
		u.Sout[i] = buf[i]
	}
}
func (u *Unit) FillIndTemper(buf []uint8) {
	len_src := len(buf)
	if(len_src > NUMBER_OUTS) {
		log.Printf("len=%d", len_src)
		len_src = NUMBER_OUTS
	}
	for i := 0; i < len_src; i++ {
		u.IndTemper[i] = buf[i]
	}
}
func (u *Unit) FillTemper(buf []uint8) {
	len_src := len(buf)
	if(len_src > NUMBER_TEMPER) {
		log.Printf("len=%d", len_src)
		len_src = NUMBER_TEMPER
	}
	for i := 0; i < len_src; i++ {
		u.Temper[i] = buf[i]
	}
}
func (u *Unit)Fill12v(s string) {
	u.U12v = s
}
func (u *Unit)FillLevelWifi(s int) {
	u.LevelWifi = s
}

func (u * Unit)PrintUnit() {
	s := u.StrUnit
	log.Printf("Unit= %s", s)
	s = utils.Hex2Str(u.Fout, 4)
	log.Printf("Fout= %s", s)
	s = utils.Hex2Str(u.Sout, 4)
	log.Printf("Sout= %s", s)
	s = utils.Hex2Str(u.Ftout, 4)
	log.Printf("Ftout= %s", s)
	s = utils.Hex2Str(u.IndTemper, 4)
	log.Printf("IndTemper= %s", s)
	s = utils.Hex2Str(u.Temper, 4)
	log.Printf("Temper= %s", s)
	log.Printf("U12v=%s", u.U12v)
	log.Printf("WiFi=%d", u.LevelWifi)

	b, err := json.Marshal(u)
	if err != nil {
		log.Printf("err %s", err)
	}
	log.Printf("b=%s", b)
}