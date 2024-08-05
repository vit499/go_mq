package http_mq

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (h *HttpServer) GetWeather(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := time.Now()
	defer func() {
		t1 := time.Since(t)
		h.logger.Info().Msgf("/api/t time: %v", t1)
	}()
	s1, err := h.unitService.GetWeather()
	if err != nil {
		//
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s1))
}
