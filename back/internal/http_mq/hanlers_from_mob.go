package http_mq

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (h *HttpServer) CmdFromMobile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := time.Now()
	defer func() {
		t1 := time.Since(t)
		h.logger.Info().Msgf("/objects/device_any_command time: %v", t1)
	}()
	objname := ps.ByName("objname")
	cmd := r.URL.Query().Get("command")
	h.logger.Info().Msgf("cmd from http, obj: %s cmd=%s ", objname, cmd)

	topic := fmt.Sprintf("ab@m.ru/%s/devrec/control", objname)
	message := cmd
	h.logger.Info().Msgf("send to mq %s %s", topic, message)
	s := []string{topic, message}
	h.hub.ChanWsToMq(s)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status", "200")
	// w.WriteHeader(http.StatusOK)
	s1 := "{ \"a\": \"ok\"}"
	w.Write([]byte(s1))
}
