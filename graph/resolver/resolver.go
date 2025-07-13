//go:generate go run github.com/99designs/gqlgen generate
package resolver

import (
	"database/sql"

	"github.com/ericgrandt/gqlgen-example/database"
)

type Resolver struct {
	db       *sql.DB
	userData database.UserData
}

func NewResolver(db *sql.DB, userData database.UserData) *Resolver {
	return &Resolver{
		db:       db,
		userData: userData,
	}
}
