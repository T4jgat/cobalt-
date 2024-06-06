package httpserver

import (
	"github.com/T4jgat/cobalt+/internal/controller/httpv1"
	"github.com/T4jgat/cobalt+/pkg/logger"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func New(rentalsController *httpv1.RentalsController, log logger.Logger) http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/rentals", rentalsController.GetAllRentals)
	r.HandlerFunc(http.MethodPost, "/rentals", rentalsController.CreateRental)
	r.HandlerFunc(http.MethodGet, "/rentals/:id", rentalsController.GetRentalByID)
	r.HandlerFunc(http.MethodDelete, "/rentals/:id", rentalsController.DeleteRentalByID)
	r.HandlerFunc(http.MethodPut, "/rentals/:id", rentalsController.UpdateRentalByID)

	r.HandlerFunc(http.MethodPost, "/rentals/user-rentals", rentalsController.GetRentalsByUserEmail)
	r.HandlerFunc(http.MethodPut, "/rentals/:id/updatestatus", rentalsController.UpdateRentalStatus)

	return r
}
