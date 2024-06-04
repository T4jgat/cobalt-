package httpserver

import (
	"github.com/T4jgat/cobalt+/internal/controller/httpv1"
	"github.com/T4jgat/cobalt+/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

func New(rentalsController *httpv1.RentalsController, log logger.Logger) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/rentals", rentalsController.CreateRental).Methods(http.MethodPost)
	r.HandleFunc("/rentals", rentalsController.GetAllRentals).Methods(http.MethodGet)
	// Add other endpoints here...

	return r
}
