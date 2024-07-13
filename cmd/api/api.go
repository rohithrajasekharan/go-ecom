package api

import (
	"database/sql"
	"net/http"

	"github.com/rohithrajasekharan/go-ecom/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(mux)
	return http.ListenAndServe(s.addr, mux)
}
