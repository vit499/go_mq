package sensor

import (
	"back/pkg/logger"
	"back/pkg/tgbot"
	"fmt"
)

/*
					"label": "Package id 0",
	        "current": 63.0,
	        "high": 105.0,
*/
type Sensor struct {
	Label   string  `json:"label"`
	Current float64 `json:"current"`
	High    float64 `json:"high"`
}

type Sensors struct {
	Sensors []Sensor `json:"sensors"`
}

type DataSensor struct {
	Tg      *tgbot.Tgbot
	logger  *logger.Logger
	sensors *Sensors
}

func NewDataSensor(tg *tgbot.Tgbot, logger *logger.Logger) *DataSensor {
	s := make([]Sensor, 0)
	return &DataSensor{tg, logger, &Sensors{Sensors: s}}
}

func (s *DataSensor) SetTemper(sensors Sensors) {
	s.sensors = &sensors

	for _, sensor := range sensors.Sensors {
		label := sensor.Label
		t := sensor.Current
		s.logger.Info().Msgf("SetTemper sensor: %s = %f", label, t)
		if t > 60 {
			str := fmt.Sprintf("N5101 sensor: %s = %f", label, t)
			s.Tg.SendMes(str)
		}
	}
}
