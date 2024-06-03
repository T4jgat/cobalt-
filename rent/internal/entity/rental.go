package entity

import "time"

type Rental struct {
	ID        string
	UserID    string
	CarID     string
	StartDate time.Time
	EndDate   time.Time
	Status    string
}
