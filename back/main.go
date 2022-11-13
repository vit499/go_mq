package main_test

import (
	//"fmt"
	"c2/pkg/utils"
	"strings"
	//"c2/internal/unit"
	"fmt"
	"log"
)


func say(ch chan [2]string) {
	for i := 0; i < 10; i++ {
		utils.D_1ms(100)
		ch <- [2]string{fmt.Sprintf("%d", i + 1), fmt.Sprintf("%d", 4)}
	}
	close(ch)
}

func FillParam(src string) {
	arrS := strings.Split(src, "&")
	for i := 0; i < len(arrS); i++ {
		log.Printf("s[%d] = %s ", i, arrS[i])
		p := strings.Split(arrS[i], "=")
		key := p[0]
		val := p[1]
		//if(key == "u") {
			log.Printf("key=%s val=%s", key, val)
		//}
	}
}

func main_test() {
	//var strW [2]string
	// strK := make(chan [2]string)
	// log.Printf("start...")
	// go say(strK)
	// log.Printf("continue...")
	// //delay.D_1s(2) 
	// //log.Printf("end")
	// for strW := range strK {
	// 	//strW = <-strK
	// 	log.Printf("w1 = %s w2 = %s ", strW[0], strW[1])
	// }

	// buf := utils.Str2Hex("35F23C54")
	// log.Printf("log main buf= % X ", buf)

	// buf2 := utils.Str2Bits("0482")
	// log.Printf("0482 -> % X ", buf2)

	// buf3 := []uint8{0x25, 0x31}
	// str3 := utils.Hex2Str(buf3, 2)
	// log.Printf("str3 = %s", str3)


	// u := unit.Unit{}
	// up := &u
	// up.Init()
	// u.FillBuf("0204", "fout")
	// u.PrintUnit()

	FillParam("dv_t1=12.5&dv_12v=12.2&dv_gsm1=0&dv_gsm2=0&tcp=0000563C&temper=800C808080808080")
}