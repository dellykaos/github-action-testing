package handler

// Handler controls request flow from client to service
type Handler struct {
}

// New initiate handler
func New() *Handler {
	return &Handler{}
}
