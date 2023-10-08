package unit

import (
	"back/pkg/utils"
	//"log"
	"strings"
)

func (u *Unit) FillBuf(topic string, src string) {
	if topic == "config/fout" {
		b := utils.Str2Hex(src)
		u.FillFout(b)
	} else if topic == "config/ftout" {
		b := utils.Str2Hex(src)
		u.FillFtout(b)
	} else if topic == "config/indtemper" {
		b := utils.Str2Hex(src)
		u.FillIndTemper(b)
	} else if topic == "status/sout" {
		b := utils.Str2Bits(src)
		u.FillSout(b)
	} else if topic == "status/param" {
		u.FillParam(src)
	} else if topic == "config/vers" {
		u.Vers = src
	}

}
func (u *Unit) FillBufEv(topic string, src string) string {
	ev := ""
	if topic == "event" {
		ev = u.RecEvent(src)
	}
	return ev
}

// dv_ev=1111E13002015&dv_time=20160313114040
func (u *Unit) RecEvent(src string) string {
	ev := ""
	arrS := strings.Split(src, "&")
	for i := 0; i < len(arrS); i++ {
		//log.Printf("s[%d] = %s ", i, arrS[i])
		p := strings.Split(arrS[i], "=")
		key := p[0]
		val := p[1]
		if key == "dv_ev" {
			ev = CheckEv(val)
		}
	}
	return ev
}

//func FillUnit(u Unit, src string)
