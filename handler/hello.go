package handler

import (
	"net/http"

	hello "github.com/dellykaos/github-action-testing"
	"github.com/julienschmidt/httprouter"
)

// Hello is used to control the flow of GET /hello endpoint
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	res := hello.Name(params.ByName("name"))

	w.Write([]byte(res))
}
