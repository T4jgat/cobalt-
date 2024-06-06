package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/T4jgat/cobalt+/internal/entity"
	"strings"
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
		rental.UserID, rental.CarID, currentTime, currentTime.Add(time.Hour*24), "PENDING")
	return err
}

//
//func (r *RentalRepo) GetAll() ([]*entity.Rental, error) {
//	rows, err := r.db.Query("SELECT id, user_id, car_id, start_date, end_date, status FROM rentals")
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	rentals := []*entity.Rental{}
//	for rows.Next() {
//		rental := &entity.Rental{}
//		err := rows.Scan(&rental.ID, &rental.UserID, &rental.CarID, &rental.StartDate, &rental.EndDate, &rental.Status)
//		if err != nil {
//			return nil, err
//		}
//		rentals = append(rentals, rental)
//	}
//
//	return rentals, nil
//}

func (r *RentalRepo) GetByID(id int) (*entity.Rental, error) {

	query := `
		SELECT id, user_id, car_id, start_date, end_date, status 
		FROM rentals 
		WHERE id = $1
		`
	var rental entity.Rental

	err := r.db.QueryRow(query, id).Scan(
		&rental.ID,
		&rental.UserID,
		&rental.CarID,
		&rental.StartDate,
		&rental.EndDate,
		&rental.Status,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &rental, nil
}

func (r *RentalRepo) DeleteByID(id int) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM ONLY rentals
		WHERE id = $1
		`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (r *RentalRepo) Update(rental *entity.Rental) error {
	query := `
		UPDATE rentals
		SET user_id = $1, car_id = $2, status = $3
		WHERE id = $4
		RETURNING id
		`

	args := []any{
		rental.UserID,
		rental.CarID,
		rental.Status,
		rental.ID,
	}

	err := r.db.QueryRow(query, args...).Scan(&rental.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return nil
		}
	}
	return nil
}

func (r *RentalRepo) CreateRental(id int) {} // TODO implement

func (r *RentalRepo) UpdateStatus(id int, status string) error {
	query := `
		UPDATE rentals
		SET status = $1
		WHERE id = $2
		RETURNING id
		`

	args := []any{
		status,
		id,
	}

	err := r.db.QueryRow(query, args...).Scan(id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return nil
		}
	}
	return nil
}

func (r *RentalRepo) GetAll(filters map[string]string, sort string) ([]*entity.Rental, error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT id, user_id, car_id, start_date, end_date, status FROM rentals WHERE 1=1")

	args := []any{}
	argID := 1

	for key, value := range filters {
		queryBuilder.WriteString(fmt.Sprintf(" AND %s = $%d", key, argID))
		args = append(args, value)
		argID++
	}

	if sort != "" {
		queryBuilder.WriteString(" ORDER BY " + sort)
	}

	query := queryBuilder.String()
	rows, err := r.db.Query(query, args...)
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
