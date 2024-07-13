package user

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/login", h.handleLogin)
}

func (h *Handler) handleLogin(w http.ResponseWriter, req *http.Request) {

}
