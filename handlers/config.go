package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ConfigureHandler configures the routes of this handler and binds handler functions to them
func ConfigureHandler(handler Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/").Handler(http.HandlerFunc(handler.Index))
	router.Methods("GET").Path("/products").
		Handler(http.HandlerFunc(handler.ProductIndex))
	router.Methods("GET").Path("/orders/{orderId}").
		Handler(http.HandlerFunc(handler.OrderShow))
	router.Methods("POST").Path("/orders").
		Handler(http.HandlerFunc(handler.OrderInsert))
	router.Methods("POST").Path("/close").Handler(http.HandlerFunc(handler.Close))
	router.Methods("GET").Path("/stats").Handler(http.HandlerFunc(handler.Stats))
	router.Methods("DELETE").Path("/orders/{orderId}").
		Handler(http.HandlerFunc(handler.OrderReverse))

	return router
}
