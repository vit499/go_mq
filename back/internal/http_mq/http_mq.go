package http_mq

import (
	//"fmt"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
