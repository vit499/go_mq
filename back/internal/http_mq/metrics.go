package http_mq

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (h *HttpServer) SetTemperN5101(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := time.Now()
	defer func() {
		t1 := time.Since(t)
		h.logger.Info().Msgf("/api/settemperN5101 time: %v", t1)
	}()

	b := []byte("{}")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error().Msgf("post io.ReadAll: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(b)
		return
	}
	err = h.sensorService.SetTemperFromN5101(body)
	if err != nil {
		h.logger.Error().Msgf("from service: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(b)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

// # HELP custom_temperature Current temperature
// # TYPE custom_temperature gauge
// custom_temperature 6.563701921747622
func (h *HttpServer) Metric(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := time.Now()
	defer func() {
		t1 := time.Since(t)
		h.logger.Info().Msgf("/metric time: %v", t1)
	}()
	s1 := h.sensorService.GetTemper()
	s2 := h.unitService.GetTemperMetric()
	s := fmt.Sprintf("%s%s", s1, s2)
	b := []byte(s)
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}
