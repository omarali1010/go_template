package domain

import (
	"context"

	"github.com/google/uuid"
)

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID uuid.UUID) (*Profile, error)
}
