package models

import "github.com/google/uuid"

type ShoppingItem struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
