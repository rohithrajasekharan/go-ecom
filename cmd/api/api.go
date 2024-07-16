package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/rohithrajasekharan/go-ecom/service/cart"
	"github.com/rohithrajasekharan/go-ecom/service/order"
	"github.com/rohithrajasekharan/go-ecom/service/products"
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

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(mux)

	productStore := products.NewStore(s.db)
	productHandler := products.NewHandler(productStore)
	productHandler.RegisterRoutes(mux)

	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(mux)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, mux)
}
