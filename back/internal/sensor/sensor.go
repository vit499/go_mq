package sensor

import (
	"back/pkg/logger"
	"back/pkg/tgbot"
	"context"
	"fmt"
	"sync"
	"time"
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
	sensors Sensors
	temper  int
	mutex   sync.Mutex
	cnt     int
}

func NewDataSensor(ctx context.Context, tg *tgbot.Tgbot, logger *logger.Logger) *DataSensor {
	s := make([]Sensor, 0)
	ds := DataSensor{tg, logger, Sensors{Sensors: s}, 0x80, sync.Mutex{}, 0}
	go ds.checkOnline(ctx)
	return &ds
}

func (ds *DataSensor) SetTemper(sensors Sensors) {
	s := make([]Sensor, 0)
	ds.sensors = Sensors{Sensors: s}
	// ds.sensors.Sensors = ds.sensors.Sensors[:0]
	// ds.sensors = &sensors

	for _, s := range sensors.Sensors {
		a := Sensor{}
		a.Label = s.Label
		a.Current = s.Current
		a.High = s.High
		ds.logger.Info().Msgf("SetTemper sensor: %s = %f", a.Label, a.Current)
		ds.sensors.Sensors = append(ds.sensors.Sensors, a)
		if a.Current > 60 {
			str := fmt.Sprintf("N5101 sensor: %s = %f", a.Label, a.Current)
			ds.Tg.SendMes(str)
		}
	}
	// for _, sensor := range sensors.Sensors {
	// 	label := sensor.Label
	// 	t := sensor.Current
	// 	ds.logger.Info().Msgf("SetTemper sensor: %s = %f", label, t)
	// 	if t > 60 {
	// 		str := fmt.Sprintf("N5101 sensor: %s = %f", label, t)
	// 		ds.Tg.SendMes(str)
	// 	}
	// }
	if len(ds.sensors.Sensors) == 0 {
		return
	}
	for _, s := range ds.sensors.Sensors {
		ds.logger.Info().Msgf(" sensor: %s = %f", s.Label, s.Current)
	}
	ds.mutex.Lock()
	ds.cnt = 0
	ds.temper = int(ds.sensors.Sensors[0].Current)
	ds.mutex.Unlock()
	ds.logger.Info().Msgf("temper = %d", ds.temper)
}

func (ds *DataSensor) GetTemper() int {
	return ds.temper
}

func (ds *DataSensor) checkOnline(ctx context.Context) {
	for {
		ticker := time.NewTicker(60 * time.Second)
		select {
		case <-ctx.Done():
			// log.Printf("ctx done checkOnline")
			return
		case <-ticker.C:
			ds.mutex.Lock()
			if ds.temper != 0x80 {
				ds.cnt++
				if ds.cnt >= 3 {
					ds.temper = 0x80
				}
			}
			ds.mutex.Unlock()
		}
	}
}
