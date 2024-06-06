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
	filters := make(map[string]string)

	color := r.URL.Query().Get("color")
	if color != "" {
		filters["color"] = color
	}

	price := r.URL.Query().Get("price")
	if price != "" {
		filters["price"] = price
	}

	sort := r.URL.Query().Get("sort")

	cars, err := c.repo.GetAll(filters, sort)
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

func (c *CatalogsController) GetCarsByModel(w http.ResponseWriter, r *http.Request) {
	model := r.URL.Query().Get("model")
	if model == "" {
		http.Error(w, "Model parameter is required", http.StatusBadRequest)
		return
	}

	cars, err := c.repo.GetByModel(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, envelope{"cars": cars}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
