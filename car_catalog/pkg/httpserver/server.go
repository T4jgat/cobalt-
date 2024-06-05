package httpserver

import (
	v1 "github.com/T4jgat/cobalt/internal/controller/http/v1"
	"github.com/T4jgat/cobalt/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

func New(catalogsController *v1.CatalogsController, log logger.Logger) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/catalogs", catalogsController.CreateCar).Methods(http.MethodPost)
	r.HandleFunc("/catalogs", catalogsController.GetAllCars).Methods(http.MethodGet)
	r.HandleFunc("/catalogs/{id:[0-9]+}", catalogsController.GetCar).Methods(http.MethodGet)
	r.HandleFunc("/catalogs/{id:[0-9]+}", catalogsController.UpdateCar).Methods(http.MethodPut)
	r.HandleFunc("/catalogs/{id:[0-9]+}", catalogsController.DeleteCar).Methods(http.MethodDelete)

	return r
}
