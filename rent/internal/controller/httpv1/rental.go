package httpv1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/T4jgat/cobalt+/helpers"
	"github.com/T4jgat/cobalt+/internal/entity"
	"github.com/T4jgat/cobalt+/internal/usecase/repo"
	"io"
	"net/http"
	"time"
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

func (c *RentalsController) GetRentalsByUserEmail(w http.ResponseWriter, r *http.Request) {

	var rentalForm struct {
		Email    string `json:"email"`
		CarModel string `json:"car_model"`
	}

	if err := json.NewDecoder(r.Body).Decode(&rentalForm); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	urlForUser := "http://localhost:4000/v1/users?rentalForm=" + rentalForm.Email
	urlForCar := "http://localhost:4001/catalogs?model=" + rentalForm.CarModel

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, urlForUser, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	req.Header.Set("Authorization", "Bearer A73NY5N3ZMAZC4FRLWLBNPOFDU")

	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	userBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		// Handle the error appropriately (e.g., return an error response)
	}

	type User struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Fname     string    `json:"fname"`
		Sname     string    `json:"sname"`
		Email     string    `json:"email"`
		Activated bool      `json:"activated"`
		UserRole  string    `json:"user_role"`
		Version   int       `json:"version"`
	}

	type UsersResponse struct {
		Users []User `json:"users"`
	}

	var usersResponse UsersResponse

	fmt.Println("al;ksdjfasdf ---- ", string(userBytes))

	err = json.Unmarshal(userBytes, &usersResponse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req, err = http.NewRequest(http.MethodGet, urlForCar, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	req.Header.Set("Authorization", "Bearer A73NY5N3ZMAZC4FRLWLBNPOFDU")

	res, err = client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer res.Body.Close()

	carBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		// Handle the error appropriately (e.g., return an error response)
	}

	type Car struct {
		ID    int    `json:"id"`
		Model string `json:"model"`
		Brand string `json:"brand"`
		Color string `json:"color"`
		Price int    `json:"price"`
	}

	type CarResponse struct {
		Cars []Car `json:"cars"`
	}

	var carResponse CarResponse

	err = json.Unmarshal(carBytes, &carResponse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var rental entity.Rental

	rental.UserID = usersResponse.Users[0].ID
	rental.CarID = carResponse.Cars[0].ID

	fmt.Println("carID: ", carResponse.Cars[0].ID)
	fmt.Println("userID:", usersResponse.Users[0].ID)

	if err = c.repo.Create(&rental); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	message := "Rent log successfully created"

	err = helpers.WriteJSON(w, http.StatusOK, envelope{"message": message}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

func (c *RentalsController) UpdateRentalStatus(w http.ResponseWriter, r *http.Request) {
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
		Status string `json:"status"`
	}

	err = helpers.ReadJSON(w, r, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rentalToUpdate.Status = input.Status

	err = c.repo.UpdateStatus(id, input.Status)
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
