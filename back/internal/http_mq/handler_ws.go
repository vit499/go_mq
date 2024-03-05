package http_mq

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *HttpServer) Ws(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.hub.ServeWs(w, r)
}
