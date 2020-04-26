//go:generate go run github.com/99designs/gqlgen

package graphql

import "github.com/zackartz/go-graphql-api/postgres"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PlacesRepo postgres.PlacesRepo
	UsersRepo  postgres.UsersRepo
}
