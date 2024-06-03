package amqprpc

import (
	"context"
	"fmt"
	"github.com/T4jgat/cobalt+/internal/usecase"
	"github.com/T4jgat/cobalt+/pkg/rabbitmq/rmq_rpc/server"
	"github.com/streadway/amqp"
)

type rentalRoutes struct {
	rentalUseCase usecase.Rental
}

func newRentalRoutes(routes map[string]server.CallHandler, usecaseRental usecase.Rental) {
	r := &rentalRoutes{usecaseRental}
	{
		routes["getRentals"] = r.
	}
}

type

func (r *rentalRoutes) getRentals() server.CallHandler {
	return func (d *amqp.Delivery) (interface{}, error) {
		rentals, err := r.rentalUseCase.Rentals(context.Background())
		if err != nil {
			return nil, fmt.Errorf("amqp_rpc - ")
		}

		res := respo
	}
}