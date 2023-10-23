package unit

import (
	"fmt"
	"strconv"
	"strings"
)

// 1111E13002015
func CheckEv(s string) string {
	res := false
	ev := strings.Split(s, "")
	nobj := strings.Join(ev[:4], "")
	cid := strings.Join(ev[4:8], "")
	// strPart := strings.Join(ev[8:10], "")
	strZone := strings.Join(ev[10:], "")
	// part, err := strconv.Atoi(strPart)
	// if err != nil {
	// 	part = 0
	// }
	zone, err := strconv.Atoi(strZone)
	if err != nil {
		zone = 0
	}
	typeMes := ""
	mes := ""

	descrObj := map[string]string{
		"0802": "Шкаф",
		"0803": "Чердак",
		"0804": "Подвал",
	}

	// log.Printf("unit= %s cid= %s zone= %d part= %d ", nobj, cid, zone, part)
	// mes = descrObj[nobj]
	if cid == "E130" {
		res = true
		mes = "Тревога"
		typeMes = "zone"
	} else if cid == "E628" {
		res = true
		typeMes = "misc"
		mes = "изменение настроек"
	} else if cid == "E702" {
		res = true
		typeMes = "misc"
		mes = "включение(перезапуск)"
	} else if cid == "E062" {
		res = true
		typeMes = "misc"
		mes = "потеря wifi"
	} else if cid == "R062" {
		res = true
		typeMes = "misc"
		mes = "восстановление wifi"
	} else if (nobj == "0804") && (cid == "E715" || cid == "R715") {
		//res = true
	} else if (nobj == "0803") && (cid == "E301") {
		res = true
		typeMes = "trbl"
		mes = "Дача неисправность 220В "
	} else if (nobj == "0803") && (cid == "R301") {
		res = true
		typeMes = "trbl"
		mes = "Дача восстановление 220В "
	} else if (nobj == "0803") && (cid == "E302") {
		res = true
		typeMes = "trbl"
		mes = "Дача неисправность акб "
	} else if (nobj == "0803") && (cid == "R302") {
		res = true
		typeMes = "trbl"
		mes = "Дача восстановление акб "
	} else if (nobj == "0803") && (cid[:2] == "R4") {
		res = true
		mes = "Дача постановка на охрану"
		typeMes = "arm"
	} else if (nobj == "0803") && (cid[:2] == "E4") {
		res = true
		mes = "Дача снятие с охраны"
		typeMes = "arm"
	} else if cid == "E701" {
		res = true
		typeMes = "misc"
		mes = "потеря связи"
	} else if cid == "R701" {
		res = true
		typeMes = "misc"
		mes = "восстановление связи"
	}

	if !res {
		return ""
	}
	// if part != 0 {
	// 	// mes = fmt.Sprintf("%s раздел %d", mes, part)
	// }

	if typeMes == "zone" {
		mes = fmt.Sprintf("%s %s зона %d", descrObj[nobj], mes, zone)
	} else if typeMes == "out" {
		mes = fmt.Sprintf("%s выход %d", mes, zone)
	} else if typeMes == "misc" {
		mes = fmt.Sprintf("%s %s ", descrObj[nobj], mes)
	}

	return mes
}
