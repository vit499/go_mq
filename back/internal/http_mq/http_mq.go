package http_mq

import (
	"back/internal/service/sensor_service"
	"back/internal/service/units_service"
	"back/internal/ws"
	"back/pkg/config"
	"back/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type HttpServer struct {
	//units  *unit.Units
	unitService   *units_service.UnitsService
	sensorService *sensor_service.SensorService
	logger        *logger.Logger
	hub           *ws.Hub
}

func GetHttpServer(ctx context.Context, unitService *units_service.UnitsService, sensorService *sensor_service.SensorService, logger *logger.Logger, hub *ws.Hub) error {
	cfg := config.Get()
	httpHost := cfg.HttpHost
	h := HttpServer{unitService, sensorService, logger, hub}

	// mux := http.NewServeMux()
	// srv := &http.Server{
	// 	Addr:    httpHost,
	// 	Handler: mux,
	// }

	//h.logger.Info().Msgf("Start HTTP %s ", httpHost)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/api/units/:ind", h.GetUnit)
	router.GET("/api/t", h.GetUnitTemper)
	router.GET("/data/2.5/weather", h.GetWeather) // http://api.openweathermap.org/data/2.5/weather?q=Kaliningrad&units=metric&APPID=0def5ea4b295f1a9d161837cb76cb667
	router.GET("/api/a", h.GetFtoutAndTemp)
	router.POST("/api/temper/n5101", h.SetTemperN5101)
	router.GET("/metrics", h.Metric)
	router.GET("/ws", h.Ws)
	router.POST("/objects/:objname/device_any_command", h.CmdFromMobile)
	srv := &http.Server{Addr: httpHost, Handler: router}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			h.logger.Error().Msgf("listen %s", err.Error())
		}
	}()

	h.logger.Info().Msgf("Start HTTP %s ", httpHost)
	<-ctx.Done()

	//h.logger.Info().Msg("ctx done http ")

	anotherCtx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()

	if err := srv.Shutdown(anotherCtx); err != nil {
		h.logger.Error().Msgf("Server shutdown: %v", err)
	}
	<-anotherCtx.Done()

	h.logger.Info().Msg("Stop HTTP ")
	return nil
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// fmt.Fprint(w, "Welcome!\n")
	s1 := "hello"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s1))
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func (h *HttpServer) GetUnit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	strInd := ps.ByName("ind")
	b, err := h.unitService.GetUnit(strInd)
	if err != nil {
		//
		return
	}
	header := w.Header()
	// header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
