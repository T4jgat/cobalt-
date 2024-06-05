package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/T4jgat/cobalt/helpers"
	"github.com/T4jgat/cobalt/internal/entity"
	"github.com/T4jgat/cobalt/internal/usecase/repo"
	"github.com/julienschmidt/httprouter"
)

type CatalogsController struct {
	repo *repo.CatalogRepo
}

func New(repo *repo.CatalogRepo) *CatalogsController {
	return &CatalogsController{repo: repo}
}

type envelope map[string]any

func (c *CatalogsController) CreateCar(w http.ResponseWriter, r *http.Request) {
	var catalog entity.Catalog

	if err := json.NewDecoder(r.Body).Decode(&catalog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.repo.Create(&catalog); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	err := helpers.WriteJSON(w, http.StatusOK, envelope{"message": "Catalog entry successfully created"}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *CatalogsController) GetAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := c.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, envelope{"cars": cars}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *CatalogsController) GetCar(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	fmt.Println(params.ByName("id"))

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	car, err := c.repo.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if car == nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, envelope{"car": car}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *CatalogsController) UpdateCar(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	var catalog entity.Catalog
	if err := json.NewDecoder(r.Body).Decode(&catalog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	catalog.ID = id

	if err := c.repo.Update(&catalog); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, envelope{"message": "Catalog entry successfully updated"}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *CatalogsController) DeleteCar(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	if err := c.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, envelope{"message": "Catalog entry successfully deleted"}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}