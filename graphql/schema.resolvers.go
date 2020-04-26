package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/zackartz/go-graphql-api/graphql/generated"
	"github.com/zackartz/go-graphql-api/graphql/model"
	"github.com/zackartz/go-graphql-api/models"
)

func (r *mutationResolver) CreatePlace(ctx context.Context, input model.NewPlace) (*models.Place, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}

	if len(input.Desc) < 3 {
		return nil, errors.New("desc not long enough")
	}

	place := &models.Place{
		Name:        input.Name,
		Description: input.Desc,
		UserID:      "1",
	}

	return r.PlacesRepo.CreatePlace(place)
}

func (r *mutationResolver) UpdatePlace(ctx context.Context, id string, input model.UpdatePlace) (*models.Place, error) {
	place, err := r.PlacesRepo.GetById(id)
	if err != nil || place == nil {
		return nil, errors.New("place does not exist")
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is not long enough")
		}
		place.Name = *input.Name
		didUpdate = true
	}

	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("desc is not long enough")
		}
		place.Name = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done")
	}

	place, err = r.PlacesRepo.Update(place)
	if err != nil {
		return nil, fmt.Errorf("error while updating place: %v", err)
	}
	return place, nil
}

func (r *mutationResolver) DeletePlace(ctx context.Context, id string) (bool, error) {
	place, err := r.PlacesRepo.GetById(id)
	if err != nil || place == nil {
		return false, errors.New("place does not exist")
	}

	err = r.PlacesRepo.Delete(place)
	if err != nil {
		return false, fmt.Errorf("error while deleting place: %v", err)
	}

	return true, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.UsersRepo.GetUsers()
}

func (r *queryResolver) Places(ctx context.Context) ([]*models.Place, error) {
	return r.PlacesRepo.GetPlaces()
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.UsersRepo.GetUserById(id)
}

func (r *userResolver) Places(ctx context.Context, obj *models.User) ([]*models.Place, error) {
	return r.PlacesRepo.GetPlacesForUser(obj)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
