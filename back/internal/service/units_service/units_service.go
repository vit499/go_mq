package units_service

import (
	"back/internal/hglob"
	"back/internal/unit"
	"errors"
	"fmt"
	"strconv"
)

type UnitsService struct {
	units *unit.Units
	hglob *hglob.Hglob
}

func NewUnitsService(us *unit.Units, hglob *hglob.Hglob) *UnitsService {
	return &UnitsService{
		units: us,
		hglob: hglob,
	}
}

func (h *UnitsService) GetUnit(strInd string) ([]byte, error) {
	ind, err := strconv.Atoi(strInd)
	if err != nil {
		//
		return nil, err
	}
	if ind >= h.units.Cnt {
		//
		return nil, nil
	}
	b, err := h.units.GetJsonUnit(ind)
	if err != nil {
		//
		return nil, err
	}
	return b, nil
}

func (h *UnitsService) GetUnitTemper() ([]byte, error) {

	temper := make([]int, 10)
	descr := []string{
		"f1",
		"",
		"",
		"f1",
		"f2",
		"",
		"outdoor",
		"f0",
		"",
	}
	for ind := 0; ind < h.units.Cnt; ind++ {
		t, _ := h.units.GetUnitTemper(ind)
		temper[ind*3] = t[0]
		temper[ind*3+1] = t[1]
		temper[ind*3+2] = t[2]
	}
	s := ""
	for ind := 0; ind < 9; ind++ {
		if temper[ind] != 0x80 {
			s = fmt.Sprintf(" %s %s = %d <br>", s, descr[ind], temper[ind])
		}
	}
	s1 := fmt.Sprintf("<html>%s</html>", s)

	return []byte(s1), nil
}

// # HELP outdoor_temperature Outdoor temperature
// # TYPE outdoor_temperature gauge
// outdoor_temperature 6.56
func (h *UnitsService) GetTemperMetric() string {
	var s1, s2, s3 string = "", "", ""
	var s10, s11, s20, s21 string = "", "", "", ""

	tempers, _ := h.units.GetUnitTemper(1)
	t := tempers[0]
	if t != 0x80 {
		s1 = "# HELP first_temperature First floor temperature\n"
		s2 = "# TYPE first_temperature gauge\n"
		s3 = fmt.Sprintf("first_temperature %d\n", t)
		s10 = fmt.Sprintf("%s%s%s", s1, s2, s3)
	}
	t = tempers[1]
	if t != 0x80 {
		s1 = "# HELP loft_temperature Loft under roof temperature\n"
		s2 = "# TYPE loft_temperature gauge\n"
		s3 = fmt.Sprintf("loft_temperature %d\n", t)
		s11 = fmt.Sprintf("%s%s%s", s1, s2, s3)
	}

	tempers, _ = h.units.GetUnitTemper(2)
	t = tempers[0]
	if t != 0x80 {
		s1 = "# HELP outdoor_temperature Outdoor temperature\n"
		s2 = "# TYPE outdoor_temperature gauge\n"
		s3 = fmt.Sprintf("outdoor_temperature %d\n", t)
		s20 = fmt.Sprintf("%s%s%s", s1, s2, s3)
	}
	t = tempers[1]
	if t != 0x80 {
		s1 = "# HELP basement_temperature Basement temperature\n"
		s2 = "# TYPE basement_temperature gauge\n"
		s3 = fmt.Sprintf("basement_temperature %d\n", t)
		s21 = fmt.Sprintf("%s%s%s", s1, s2, s3)
	}

	s := fmt.Sprintf("%s%s%s%s", s10, s11, s20, s21)
	return s
}

// func (h *UnitsService) FormJsonToWs(user string) []string {
// 	for ind := 0; ind < h.units.Cnt; ind++ {

// 		topic := h.units.Up[ind].StrUnit
// 		b, err := h.units.GetJsonUnit(ind)
// 		if err != nil {
// 			//
// 			continue
// 		}
// 		s := []string{user, topic, string(b)}
// 		// log.Printf("FormJsonToWs to <-HtToWs %s", s[0])
// 		// h.hglob.HtToWs <- s
// 		return s
// 	}
// }

func (h *UnitsService) GetCountUnits() int {
	return h.units.Cnt
}

func (h *UnitsService) GetUnitByInd(ind int, user string) ([]string, error) {
	topic := h.units.Up[ind].StrUnit
	b, err := h.units.GetJsonUnit(ind)
	if err != nil {
		//
		return nil, err
	}
	s := []string{user, topic, string(b)}
	return s, nil
}

func (h *UnitsService) GetUnitByName(strUnit string, user string) ([]string, error) {
	for i := 0; i < h.units.Cnt; i++ {
		if h.units.Up[i].StrUnit == strUnit {
			return h.GetUnitByInd(i, user)
		}
	}
	return nil, errors.New("not found")
}

func (h *UnitsService) CheckPass(pass string) bool {
	return h.units.CheckPass(pass)
}
