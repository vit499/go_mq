package http_mq

import (
	//service "back/internal/service/units_service"
	"back/internal/service/units_service"
	"back/internal/ws"
	"back/pkg/config"
	"back/pkg/logger"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type HttpServer struct {
	//units  *unit.Units
	service *service.UnitsService
	logger  *logger.Logger
	hub     *ws.Hub
}

func GetHttpServer(ctx context.Context, service *service.UnitsService, logger *logger.Logger, hub *ws.Hub) error {
	cfg := config.Get()
	httpHost := cfg.HttpHost
	h := HttpServer{service, logger, hub}

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
	router.GET("/ws", h.Ws)
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

// func (h *HttpServer) StartHttp(addr string) error {
// 	router := httprouter.New()
// 	router.GET("/", Index)
// 	router.GET("/hello/:name", Hello)
// 	router.GET("/api/units/:ind", h.GetUnit)
// 	router.GET("/api/t", h.GetUnitTemper)
// 	router.GET("/ws", h.Ws)

// 	err := http.ListenAndServe(addr, router)
// 	return err
// }

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
func (h *HttpServer) Ws(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//ws.ServeWs(h.hub, w, r)
	//h.logger.Info().Msgf("req ws %v", ps)
	h.hub.ServeWs(w, r)
	//h.service.FormJsonToWs("ab@m.ru")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func (h *HttpServer) GetUnit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	strInd := ps.ByName("ind")
	b, err := h.service.GetUnit(strInd)
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

func (h *HttpServer) GetUnitTemper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := time.Now()
	defer func() {
		t1 := time.Since(t)
		h.logger.Info().Msgf("/api/t time: %v", t1)
	}()
	b, err := h.service.GetUnitTemper()
	if err != nil {
		//
		return
	}

	//s := fmt.Sprintf("outdoor=%d, floor0=%d", temper[0], temper[1])
	// header := w.Header()
	// header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}
