package unit

import (
	"c2/pkg/utils"
	"log"
)

const NUMBER_OUTS = 32
const NUMBER_TEMPER = 32

type Unit struct {
	Fout []uint8       // function of outs
	Sout []uint8       // status outs
	Ftout []uint8      // temper for ON-OFF
	IndTemper []uint8  // index of tempers for out
	Temper []uint8     // value tempers
	U12v string
	LevelGsm int
	LevelWifi int
}

func (u *Unit) Init() {
	//u.Fout := make([]uint8, 32)
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
	for i := 0; i < len_src; i++ {
		u.Fout[i] = buf[i]
	}
}
func (u *Unit) FillFtout(buf []uint8) {
	len_src := len(buf)
	for i := 0; i < len_src; i++ {
		u.Ftout[i] = buf[i]
	}
}
func (u *Unit) FillSout(buf []uint8) {
	len_src := len(buf)
	for i := 0; i < len_src; i++ {
		u.Sout[i] = buf[i]
	}
}
func (u *Unit) FillIndTemper(buf []uint8) {
	len_src := len(buf)
	for i := 0; i < len_src; i++ {
		u.IndTemper[i] = buf[i]
	}
}
func (u *Unit) FillTemper(buf []uint8) {
	len_src := len(buf)
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
	s := utils.Hex2Str(u.Fout, NUMBER_OUTS)
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

}