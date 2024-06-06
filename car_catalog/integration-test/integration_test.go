package integration_test

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"os"
	"testing"
)

var (
	postgresContainer testcontainers.Container
	rabbitMQContainer testcontainers.Container
)

func TestMain(m *testing.M) {
	// Stop containers after test execution
	ctx := context.Background()
	os.Exit(func() int {
		code := m.Run()
		if code == 0 {
			// Stop RabbitMQ container
			err := rabbitMQContainer.Terminate(ctx)
			if err != nil {
				fmt.Println("Error stopping RabbitMQ container:", err)
			}

			// Stop PostgreSQL container
			err = postgresContainer.Terminate(ctx)
			if err != nil {
				fmt.Println("Error stopping PostgreSQL container:", err)
			}
		}
		return code
	}())
}
