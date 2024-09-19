package resolvers

import "github.com/BRAVO68WEB/collaborate-with-me/packages/api/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repositories *repository.Repository
}
