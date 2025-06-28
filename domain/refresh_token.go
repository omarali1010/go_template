package domain

import (
	"context"

	"github.com/google/uuid"
)

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUsecase interface {
	GetUserByID(c context.Context, id uuid.UUID) (User, error)
	CreateAccessToken(user *User) (accessToken string, err error)
	CreateRefreshToken(user *User) (refreshToken string, err error)
	GetRefreshAndAccessToken(oldRefreshToken string, c context.Context) (refreshTokenResopnse RefreshTokenResponse, err error)
	ExtractIDFromToken(requestToken string, secret string) (uuid.UUID, error)
}
