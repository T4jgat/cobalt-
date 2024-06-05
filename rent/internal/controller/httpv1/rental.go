package httpv1

import (
	"encoding/json"
	"errors"
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
	//
	//requestURL := "http://localhost"
	//res, err := http.Get(requestURL)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	//var rental entity.Rental = helpers.ReadJSON()

	//fmt.Println("Response: %w")

	if err := c.repo.Create(&rental); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	message := "Rent log successfully created"

	err := helpers.WriteJSON(w, http.StatusOK, envelope{"message": message}, nil)
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

	err = helpers.WriteJSON(w, http.StatusOK, envelope{"rentals": rentals}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

func (c *RentalsController) GetRentalByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIDPAram(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	rental, err := c.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, envelope{"rental": rental}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (c *RentalsController) DeleteRentalByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIDPAram(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	err = c.repo.DeleteByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, envelope{"message": "rental successfully deleted"}, nil)
}

func (c *RentalsController) UpdateRentalByID(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIDPAram(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	rentalToUpdate, err := c.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var input struct {
		UserID int    `json:"user_id"`
		CarID  int    `json:"car_id"`
		Status string `json:"status"`
	}

	err = helpers.ReadJSON(w, r, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rentalToUpdate.UserID = input.UserID
	rentalToUpdate.CarID = input.CarID
	rentalToUpdate.Status = input.Status

	err = c.repo.Update(rentalToUpdate)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrEditConflict):
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, envelope{"rental": rentalToUpdate}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
