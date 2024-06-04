package httpv1

import (
	"encoding/json"
	"github.com/T4jgat/cobalt+/helpers"
	"github.com/T4jgat/cobalt+/internal/entity"
	"github.com/T4jgat/cobalt+/internal/usecase/repo"
	"net/http"
)

type RentalsController struct {
	repo *repo.RentalRepo
}

func New(repo *repo.RentalRepo) *RentalsController {
	return &RentalsController{repo: repo}
}

type envelope map[string]any

func (c *RentalsController) CreateRental(w http.ResponseWriter, r *http.Request) {
	var rental entity.Rental

	if err := json.NewDecoder(r.Body).Decode(&rental); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.repo.Create(&rental); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	err := helpers.WriteJSON(w, http.StatusOK, envelope{"message": "Rental log successfully deleted"}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}

func (c *RentalsController) GetAllRentals(w http.ResponseWriter, r *http.Request) {
	rentals, err := c.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(rentals)
}
