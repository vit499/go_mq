package sensor_service

import (
	"back/internal/sensor"
	"back/pkg/logger"
	"encoding/json"
	"errors"
	"fmt"
)

type SensorService struct {
	dataSensor *sensor.DataSensor
	logger     *logger.Logger
}

func NewSensorService(dataSensor *sensor.DataSensor, logger *logger.Logger) *SensorService {
	return &SensorService{dataSensor, logger}
}
func (s *SensorService) SetTemperFromN5101(r []byte) error {

	// s.logger.Info().Msgf("r: %s", string(r))
	var sensors sensor.Sensors
	err := json.Unmarshal(r, &sensors)
	if err != nil {
		return errors.New("invalid request body")
	}
	// s.logger.Info().Msgf("SetTemperFromN5101 sensors: %v", sensors)

	s.dataSensor.SetTemper(sensors)
	return nil
}

// # HELP custom_temperature Current temperature
// # TYPE custom_temperature gauge
// custom_temperature 6.563701921747622
func (sens *SensorService) GetTemper() string {
	var s1, s2, s3 string = "", "", ""
	t := sens.dataSensor.GetTemper()

	if t != 0x80 {
		s1 = "# HELP hm_n5101_temperature Current n5101 temperature\n"
		s2 = "# TYPE hm_n5101_temperature gauge\n"
		s3 = fmt.Sprintf("hm_n5101_temperature %d\n", t)
	}

	s := fmt.Sprintf("%s%s%s", s1, s2, s3)
	return s
}
