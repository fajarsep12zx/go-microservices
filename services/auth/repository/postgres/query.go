package postgres

import (
	"go-microservices/core/connection"
)

// Repository ...
type Repository struct {
	Connection connection.Connection
}
