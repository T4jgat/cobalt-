package v1

import (
	"github.com/T4jgat/cobalt/pkg/logger"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewRouter(catalogsController *CatalogsController, log logger.Logger) http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodPost, "/catalogs", catalogsController.CreateCar)
	r.HandlerFunc(http.MethodGet, "/catalogs", catalogsController.GetAllCars)
	r.HandlerFunc(http.MethodGet, "/catalogs/:id", catalogsController.GetCar)
	r.HandlerFunc(http.MethodPut, "/catalogs/:id", catalogsController.UpdateCar)
	r.HandlerFunc(http.MethodDelete, "/catalogs/:id", catalogsController.DeleteCar)

	return r
}
