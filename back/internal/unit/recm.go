package unit

import (
	"back/pkg/utils"
	//"log"
	"strings"
)

func (u *Unit)FillBuf(topic string, src string) {
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
	}
}

// dv_t1=12.5&dv_12v=12.2&dv_gsm1=0&dv_gsm2=0&tcp=0000563C&temper=800C808080808080
func (u *Unit)FillParam(src string) {
	arrS := strings.Split(src, "&")
	for i := 0; i < len(arrS); i++ {
		//log.Printf("s[%d] = %s ", i, arrS[i])
		p := strings.Split(arrS[i], "=")
		key := p[0]
		val := p[1]
		if(key == "dv_12v") {
		  u.Fill12v(val)
		} else if(key == "tcp") {
		  u.FillTcp(val)
		} else if(key == "temper") {
			b := utils.Str2Hex(val)
		  u.FillTemper(b)
		}
	}
}

func (u *Unit)FillTcp(s string) {
	b := utils.Str2Hex(s)
	u.FillLevelWifi(int(b[3]))
}

//func FillUnit(u Unit, src string)