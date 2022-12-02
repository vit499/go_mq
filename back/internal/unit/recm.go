package unit

import (
	"back/pkg/utils"
	//"log"
	"strings"
)

func (u *Unit) FillBuf(topic string, src string) string {
	ev := ""
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
		b := utils.Str2Hex(src)
		u.FillSout(b)
	} else if topic == "status/param" {
		u.FillParam(src)
	} else if topic == "status/cp" {
		//u.FillParam(src)
	} else if topic == "event" {
		ev = u.RecEvent(src)
	}
	return ev
}

// dv_t1=12.5&dv_12v=12.2&dv_gsm1=0&dv_gsm2=0&tcp=0000563C&temper=800C808080808080
func (u *Unit) FillParam(src string) {
	arrS := strings.Split(src, "&")
	for i := 0; i < len(arrS); i++ {
		//log.Printf("s[%d] = %s ", i, arrS[i])
		p := strings.Split(arrS[i], "=")
		key := p[0]
		val := p[1]
		if key == "dv_12v" {
			u.Fill12v(val)
		} else if key == "tcp" {
			u.FillTcp(val)
		} else if key == "temper" {
			b := utils.Str2Hex(val)
			u.FillTemper(b)
		}
	}
}

// 1111E13002015
func (u *Unit) CheckEv(s string) string {
	res := true
	ev := strings.Split(s, "")
	nobj := strings.Join(ev[:4], "")
	cid := strings.Join(ev[4:8], "")
	// part := strings.Join(ev[8:10], "")
	// zone := strings.Join(ev[10:], "")

	//log.Printf("unit= %s cid= %s zone= %s part= %s ", nobj, cid, zone, part)
	// if cid == "E715" {
	// 	res = false
	// }
	if nobj == "0804" {
		return s
	}
	switch cid {
	case "E715":
		res = false
	case "R715":
		res = false
	}
	if !res {
		return ""
	}
	return s
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
			ev = u.CheckEv(val)
		}
	}
	return ev
}

func (u *Unit) FillTcp(s string) {
	b := utils.Str2Hex(s)
	u.FillLevelWifi(int(b[3]))
}

//func FillUnit(u Unit, src string)
