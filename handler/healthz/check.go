package healthz

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handler controls request flow from client to service
type Handler struct {
}

// New initiate healthz handler
func New() *Handler {
	return &Handler{}
}

// Healthz is used to control the flow of GET /healthz endpoint
func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("ok"))
}
