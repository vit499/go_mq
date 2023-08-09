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

func (s *SensorService) GetTemper() []byte {
	t := s.dataSensor.GetTemper()
	b := []byte(fmt.Sprintf("{temper: %d}", t))
	return b
}
