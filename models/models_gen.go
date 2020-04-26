// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *User      `json:"user"`
}

type AuthToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewPlace struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type PlaceFilter struct {
	Name *string `json:"name"`
}

type RegisterInput struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
}

type UpdatePlace struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}