package http_mq

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (h *HttpServer) GetUnitTemper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := time.Now()
	defer func() {
		t1 := time.Since(t)
		h.logger.Info().Msgf("/api/t time: %v", t1)
	}()
	s1, err := h.unitService.GetUnitTemper()
	if err != nil {
		//
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s1))
}

func (h *HttpServer) GetFtoutAndTemp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := time.Now()
	defer func() {
		t1 := time.Since(t)
		h.logger.Info().Msgf("/api/a time: %v", t1)
	}()
	s1, err := h.unitService.GetFtoutAndTemp()
	if err != nil {
		//
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status", "200")
	// w.WriteHeader(http.StatusOK)
	w.Write(s1)
}
