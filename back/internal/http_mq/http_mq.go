package http_mq

import (
	"back/internal/unit"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type HttpServer struct {
	units *unit.Units
}

func GetHttpServer(us *unit.Units) *HttpServer {
	h := HttpServer{units: us}
	return &h
}

// func (h *HttpServer) StartHttp(addr string) error {
// 	gin.DisableConsoleColor()
// 	router := gin.Default()
// 	router.Use(gin.Logger())

// 	router.GET("/", h.Hello)
// 	router.GET("/api/units/:ind", h.GetUnit)
// 	router.GET("/api/units", h.GetUnits)
// 	router.GET("/user/:name", h.HelloUser)
// 	err := router.Run(addr)
// 	return err
// }

// func (h *HttpServer) Hello(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "ok",
// 	})
// }
// func (h *HttpServer) HelloUser(c *gin.Context) {
// 	name := c.Param("name")
// 	c.String(http.StatusOK, "hello %s ", name)
// }

// func (h *HttpServer) GetUnit(c *gin.Context) {
// 	strInd := c.Param("ind")
// 	ind, err := strconv.Atoi(strInd)
// 	if err != nil {
// 		//
// 		return
// 	}
// 	if ind >= h.units.Cnt {
// 		//
// 		return
// 	}
// 	b, err := h.units.GetJsonUnit(ind)
// 	if err != nil {
// 		//
// 		return
// 	}
// 	c.JSON(http.StatusOK, string(b))
// }
// func (h *HttpServer) GetUnits(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "ok",
// 	})
// }

func (h *HttpServer) StartHttp(addr string) error {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/api/units/:ind", h.GetUnit)
	router.GET("/api/t", h.GetUnitTemper)

	err := http.ListenAndServe(addr, router)
	return err
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
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
