package usecase

import (
	"context"
	"github.com/T4jgat/cobalt+/internal/entity"
)

type (
	Rental interface {
		Rentals(context.Context) ([]entity.Rental, error)
	}

	RentalRepo interface {
	}

	RentalWebApi interface {
	}
)
