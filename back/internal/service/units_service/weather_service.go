package units_service

import (
	// "encoding/json"
	"fmt"
)

// {
//   "coord": {
//     "lon": 20.511,
//     "lat": 54.7065
//   },
//   "weather": [
//     {
//       "id": 804,
//       "main": "Clouds",
//       "description": "overcast clouds",
//       "icon": "04n"
//     }
//   ],
//   "base": "stations",
//   "main": {
//     "temp": 16.03,
//     "feels_like": 15.93,
//     "temp_min": 16.03,
//     "temp_max": 16.03,
//     "pressure": 1009,
//     "humidity": 86,
//     "sea_level": 1009,
//     "grnd_level": 1006
//   },
//   "visibility": 10000,
//   "wind": {
//     "speed": 0.99,
//     "deg": 272,
//     "gust": 1.4
//   },
//   "clouds": {
//     "all": 100
//   },
//   "dt": 1722803137,
//   "sys": {
//     "type": 1,
//     "id": 8934,
//     "country": "RU",
//     "sunrise": 1722740038,
//     "sunset": 1722796443
//   },
//   "timezone": 7200,
//   "id": 554234,
//   "name": "Kaliningrad",
//   "cod": 200
// }

// {
//   "weather": [
//     {
//       "id": 804,
//       "main": "Clouds",
//       "description": "overcast clouds",
//       "icon": "04n"
//     }
//   ],
//   "main": {
//     "temp": 16.03
//   }
// }

// {"weather": [{"icon": "04n"}],"main": {"temp": 15}}

func (h *UnitsService) GetWeather() (string, error) {

	// temper := make([]int, 10)

	// descr := "{\"weather\": [{\"icon\": \"04n\"}],\"main\": {\"temp\": 16.03}}"
	t, _ := h.units.GetUnitTemper(2)
	temp := t[0]
	s1 := fmt.Sprintf("{\"weather\": [{\"icon\": \"04n\"}],\"main\": {\"temp\": %d}}", temp)

	// s1 := fmt.Sprintf("<html>%s</html>", s)

	return s1, nil
}
