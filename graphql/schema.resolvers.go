package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/crossplatform-mapdb/graphql-api/middleware"
	"log"

	"github.com/crossplatform-mapdb/graphql-api/graphql/generated"
	"github.com/crossplatform-mapdb/graphql-api/models"
)

var (
	ErrBadCredentials  = errors.New("email/password combination didn't work")
	ErrUnauthenticated = errors.New("unauthenticated")
)

func (r *mutationResolver) Register(ctx context.Context, input *models.RegisterInput) (*models.AuthResponse, error) {
	_, err := r.UsersRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email already in use")
	}

	_, err = r.UsersRepo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("username already taken")
	}

	user := &models.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}
	// TODO: send verification code

	tx, err := r.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("error creating a transaction: %v", err)
		return nil, errors.New("something went wrong")
	}
	defer tx.Rollback()
	if _, err := r.UsersRepo.CreateUser(tx, user); err != nil {
		log.Printf("error creating a transaction: %v", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("error while committing: %v", err)
		return nil, err
	}

	token, err := user.GenToken()
	if err != nil {
		log.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input *models.LoginInput) (*models.AuthResponse, error) {
	user, err := r.UsersRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, ErrBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, ErrBadCredentials
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, ErrBadCredentials
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *mutationResolver) CreatePlace(ctx context.Context, input models.NewPlace) (*models.Place, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}

	if len(input.Desc) < 3 {
		return nil, errors.New("desc not long enough")
	}

	place := &models.Place{
		Name:        input.Name,
		Description: input.Desc,
		UserID:      currentUser.ID,
	}

	return r.PlacesRepo.CreatePlace(place)
}

func (r *mutationResolver) UpdatePlace(ctx context.Context, id string, input models.UpdatePlace) (*models.Place, error) {
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

func (r *queryResolver) Places(ctx context.Context, filter *models.PlaceFilter, limit *int, offset *int) ([]*models.Place, error) {
	return r.PlacesRepo.GetPlaces(filter, limit, offset)
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
