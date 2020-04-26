package models

type Place struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}
