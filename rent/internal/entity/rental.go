package entity

import "time"

type Rental struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	CarID     int       `json:"car_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
}
