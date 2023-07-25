package service

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
