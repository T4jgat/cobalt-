package repo

import (
	"database/sql"
	"github.com/T4jgat/cobalt+/internal/entity"
	"time"
)

type RentalRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *RentalRepo {
	return &RentalRepo{db: db}
}

func (r *RentalRepo) Create(rental *entity.Rental) error {

	currentTime := time.Now()

	_, err := r.db.Exec("INSERT INTO rentals (user_id, car_id, start_date, end_date, status) VALUES ($1, $2, $3, $4, $5)",
		rental.UserID, rental.CarID, currentTime, currentTime.Add(time.Hour*24), rental.Status)
	return err
}

func (r *RentalRepo) GetAll() ([]*entity.Rental, error) {
	rows, err := r.db.Query("SELECT id, user_id, car_id, start_date, end_date, status FROM rentals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rentals := []*entity.Rental{}
	for rows.Next() {
		rental := &entity.Rental{}
		err := rows.Scan(&rental.ID, &rental.UserID, &rental.CarID, &rental.StartDate, &rental.EndDate, &rental.Status)
		if err != nil {
			return nil, err
		}
		rentals = append(rentals, rental)
	}

	return rentals, nil
}
