package amqprpc

import (
	"github.com/T4jgat/cobalt+/internal/usecase"
	"github.com/T4jgat/cobalt+/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Rental) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
