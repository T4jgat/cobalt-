package usecase

import (
	"context"
	"github.com/T4jgat/cobalt+/internal/entity"
)

type RentUseCase interface {
	CreateRental(ctx context.Context, rental entity.Rental) error
	GetRentals(ctx context.Context, userID string) ([]entity.Rental, error)
	GetRentalByID(ctx context.Context, rentalID string) (*entity.Rental, error)
	UpdateRental(ctx context.Context, rental entity.Rental) error
	DeleteRental(ctx context.Context, rentalID string) error
}
