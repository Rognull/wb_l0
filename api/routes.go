package api

import (
	"l0/internal/handler"
	"github.com/gorilla/mux"
)

func CreateRoutes(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()  
	r.HandleFunc("/order/{order_uid}",h.GetOrder).Methods("GET")
	r.HandleFunc("/newOrder",h.NewOrder).Methods("POST")
	r.NotFoundHandler = r.NewRoute().HandlerFunc(handler.NotFound).GetHandler() 
	return r
}