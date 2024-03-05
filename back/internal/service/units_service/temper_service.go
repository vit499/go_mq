package units_service

import (
	"encoding/json"
	"fmt"
)

func (h *UnitsService) GetUnitTemper() (string, error) {

	temper := make([]int, 10)
	descr := []string{
		"f11",
		"f12",
		"f13",
		"f21",
		"f22",
		"f23",
		"outdoor",
		"f02",
		"f03",
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
	if s == "" {
		for ind := 0; ind < 9; ind++ {
			s = fmt.Sprintf(" %s %s = %d <br>", s, descr[ind], temper[ind])
		}
	}
	s1 := fmt.Sprintf("<html>%s</html>", s)

	return s1, nil
}

type FtoutAndTemp struct {
	StrUnit   string
	NameOut   string
	Output    int
	Ftout     int
	IndTemper int
	Temper    int
}

func (h *UnitsService) GetFtoutAndTemp() ([]byte, error) {

	nameOut := []string{
		"Кот", "Шкаф", "", "",
		"", "", "Чердак", "",
		"Подвал", "", "", "",
	}

	ft := make([]FtoutAndTemp, 0)

	for ind := 0; ind < h.units.Cnt; ind++ {
		name := h.units.Up[ind].StrUnit
		fout := h.units.Up[ind].Fout

		for o := 0; o < len(fout); o++ {
			if fout[o] == 7 { // 7 - temperature
				ftout := h.units.Up[ind].Ftout[o]
				indTemper := h.units.Up[ind].IndTemper[o]
				temper := h.units.Up[ind].Temper[indTemper]
				ft = append(ft, FtoutAndTemp{name, nameOut[ind*4+o], o, ftout, indTemper, temper})

			}
		}
	}

	s1, _ := json.Marshal(ft)
	h.logger.Info().Msgf("s1=%s", s1)

	return s1, nil
}

// host/objects/0802/device_any_command?command=setout1,25,0
