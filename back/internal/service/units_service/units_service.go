package units_service

import (
	"back/internal/hglob"
	"back/internal/unit"
	"back/pkg/logger"
	"errors"
	"fmt"
	"strconv"
)

type UnitsService struct {
	units  *unit.Units
	hglob  *hglob.Hglob
	logger *logger.Logger
}

func NewUnitsService(us *unit.Units, hglob *hglob.Hglob, logger *logger.Logger) *UnitsService {
	return &UnitsService{
		units:  us,
		hglob:  hglob,
		logger: logger,
	}
}

func (h *UnitsService) GetUnit(strInd string) ([]byte, error) {
	ind, err := strconv.Atoi(strInd)
	if err != nil {
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

// # HELP outdoor_temperature Outdoor temperature
// # TYPE outdoor_temperature gauge
// outdoor_temperature 6.56
func (h *UnitsService) GetTemperMetric() string {
	var s1, s2, s3 string = "", "", ""
	var s00, s01, s10, s11, s20, s21, s30, s40, s41 string = "", "", "", "", "", "", "", "", ""

	tempers, _ := h.units.GetUnitTemper(0)
	t := tempers[0]
	if t != 0x80 {
		s1 = "# HELP first_temperature First floor temperature\n"
		s2 = "# TYPE first_temperature gauge\n"
		s3 = fmt.Sprintf("first_temperature %d\n", t)
		s00 = fmt.Sprintf("%s%s%s", s1, s2, s3)
	}
	t = tempers[1]
	if t != 0x80 {
		s1 = "# HELP loft_temperature Cats room temperature\n"
		s2 = "# TYPE cats_temperature gauge\n"
		s3 = fmt.Sprintf("cats_temperature %d\n", t)
		s01 = fmt.Sprintf("%s%s%s", s1, s2, s3)
	}

	tempers, _ = h.units.GetUnitTemper(1)
	t = tempers[0]
	if t != 0x80 {
		s1 = "# HELP first_temperature First floor temperature1\n"
		s2 = "# TYPE first_temperature1 gauge\n"
		s3 = fmt.Sprintf("first_temperature1 %d\n", t)
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

	tempers, _ = h.units.GetUnitTemper(3)
	t = tempers[2]
	if t != 0x80 {
		s1 = "# HELP utro4x4_temperature Utro4x4 temperature\n"
		s2 = "# TYPE utro4x4_temperature gauge\n"
		s3 = fmt.Sprintf("utro4x4_temperature %d\n", t)
		s30 = fmt.Sprintf("%s%s%s", s1, s2, s3)
	}

	tempers, _ = h.units.GetUnitTemper(4)
	t = tempers[0]
	if t != 0x80 {
		s1 = "# HELP utro3x3_temperature Utro3x3 temperature\n"
		s2 = "# TYPE utro3x3_temperature gauge\n"
		s3 = fmt.Sprintf("utro3x3_temperature %d\n", t)
		s40 = fmt.Sprintf("%s%s%s", s1, s2, s3)
	}
	t = tempers[1]
	if t != 0x80 {
		s1 = "# HELP utro3x32_temperature Utro3x32 temperature\n"
		s2 = "# TYPE utro3x32_temperature gauge\n"
		s3 = fmt.Sprintf("utro3x32_temperature %d\n", t)
		s41 = fmt.Sprintf("%s%s%s", s1, s2, s3)
	}

	s := fmt.Sprintf("%s%s%s%s%s%s%s%s%s", s00, s01, s10, s11, s20, s21, s30, s40, s41)
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
