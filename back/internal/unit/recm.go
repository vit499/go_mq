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
		b := utils.Str2Bits(src)
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

// 1111E13002015
func (u *Unit) CheckEv(s string) string {
	res := false
	ev := strings.Split(s, "")
	nobj := strings.Join(ev[:4], "")
	cid := strings.Join(ev[4:8], "")
	// part := strings.Join(ev[8:10], "")
	// zone := strings.Join(ev[10:], "")

	//log.Printf("unit= %s cid= %s zone= %s part= %s ", nobj, cid, zone, part)
	if cid == "E130" {
		res = true
	} else if cid == "E628" {
		res = true
	} else if cid == "E702" {
		res = true
	} else if cid == "E062" || cid == "R062" {
		res = true
	} else if (nobj == "0804") && (cid == "E715" || cid == "R715") {
		//res = true
	} else if (nobj == "0803") && (cid == "E301" || cid == "R301") {
		res = true
	} else if cid == "E409" || cid == "R409" {
		res = true
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

//func FillUnit(u Unit, src string)
