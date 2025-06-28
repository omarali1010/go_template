package domain

import "github.com/google/uuid"

type AccessTokenCreator interface {
	CreateAccessToken(user *User) (string, error)
}

type RefreshTokenCreator interface {
	CreateRefreshToken(user *User) (string, error)
}

type AccessTokenParser interface {
	GetIDFromAccessToken(token string) (uuid.UUID, error)
}

type RefreshTokenParser interface {
	GetIDFromRefreshToken(token string) (uuid.UUID, error)
}
