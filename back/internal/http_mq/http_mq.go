package http_mq

import (
	"back/internal/unit"
	"back/internal/ws"
	"back/pkg/config"
	"back/pkg/logger"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type HttpServer struct {
	units  *unit.Units
	logger *logger.Logger
	hub    *ws.Hub
}

func GetHttpServer(us *unit.Units, logger *logger.Logger, hub *ws.Hub) (*HttpServer, error) {
	cfg := config.Get()
	httpHost := cfg.HttpHost
	h := HttpServer{units: us, logger: logger, hub: hub}

	h.logger.Info().Msgf("Start HTTP %s ", httpHost)
	err := h.StartHttp(httpHost)
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (h *HttpServer) StartHttp(addr string) error {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/api/units/:ind", h.GetUnit)
	router.GET("/api/t", h.GetUnitTemper)
	router.GET("/ws", h.Ws)

	err := http.ListenAndServe(addr, router)
	return err
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
func (h *HttpServer) Ws(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//fmt.Fprint(w, "Welcome!\n")
	//ws.ServeWs(h.hub, w, r)
	h.hub.ServeWs(w, r)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func (h *HttpServer) GetUnit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	strInd := ps.ByName("ind")
	ind, err := strconv.Atoi(strInd)
	if err != nil {
		//
		return
	}
	if ind >= h.units.Cnt {
		//
		return
	}
	b, err := h.units.GetJsonUnit(ind)
	if err != nil {
		//
		return
	}

	// log.Printf("cors ?")
	// if r.Header.Get("Access-Control-Request-Method") != "" {
	// 	log.Printf("cors ?")
	// 	// Set CORS headers
	// 	header := w.Header()
	// 	header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
	// 	header.Set("Access-Control-Allow-Origin", "*")
	// }
	header := w.Header()
	// header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h *HttpServer) GetUnitTemper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	temper := make([]int, 10)
	descr := []string{
		"f1",
		"",
		"",
		"f1",
		"f2",
		"",
		"outdoor",
		"f0",
		"",
	}
	for ind := 0; ind < h.units.Cnt; ind++ {
		t, _ := h.units.GetUnitTemper(ind)
		temper[ind*3] = t[0]
		temper[ind*3+1] = t[1]
		temper[ind*3+2] = t[2]
	}
	s := ""
	for ind := 0; ind < 9; ind++ {
		if temper[ind] != 0x80 {
			s = fmt.Sprintf(" %s %s = %d <br>", s, descr[ind], temper[ind])
		}
	}
	s1 := fmt.Sprintf("<html>%s</html>", s)

	//s := fmt.Sprintf("outdoor=%d, floor0=%d", temper[0], temper[1])
	// header := w.Header()
	// header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s1))
}
