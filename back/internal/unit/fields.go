package unit

import (
	"back/pkg/utils"
	"strings"
)

func (u *Unit) FillTcp(s string) {
	b := utils.Str2Hex(s)
	u.FillLevelWifi(int(b[3]))
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

func (u *Unit) Fill12v(s string) {
	u.U12v = s
}
func (u *Unit) FillLevelWifi(s int) {
	u.LevelWifi = s
}
