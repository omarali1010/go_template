package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/omaraliali1010/go_template/domain"
)

type refreshTokenUserCase struct {
	userRepository      domain.UserRepository
	accesTokenCreator   domain.AccessTokenCreator
	refreshTokenCreator domain.RefreshTokenCreator
	refreshTokenParser  domain.RefreshTokenParser
	contextTimeout      time.Duration
}

func NewRefreshTokenUserCase(userRepository domain.UserRepository,
	accesTokenCreator domain.AccessTokenCreator,
	refreshTokenCreator domain.RefreshTokenCreator,
	refreshTokenParser domain.RefreshTokenParser,
	timeout time.Duration) domain.RefreshTokenUsecase {

	return &refreshTokenUserCase{
		userRepository:      userRepository,
		accesTokenCreator:   accesTokenCreator,
		refreshTokenCreator: refreshTokenCreator,
		refreshTokenParser:  refreshTokenParser,
		contextTimeout:      timeout,
	}
}

// GetRefreshAndAccessToken implements domain.RefreshTokenUsecase.
func (r *refreshTokenUserCase) GetRefreshAndAccessToken(oldRefreshToken string, c context.Context) (refreshTokenResopnse domain.RefreshTokenResponse, err error) {
	id, err := r.refreshTokenParser.GetIDFromRefreshToken(oldRefreshToken)
	if err != nil {
		return domain.RefreshTokenResponse{}, err
	}

	user, err := r.GetUserByID(c, id)
	if err != nil {
		return domain.RefreshTokenResponse{}, err
	}
	accessToken, err := r.CreateAccessToken(&user)
	if err != nil {
		return domain.RefreshTokenResponse{}, err
	}
	refreshToken, err := r.CreateRefreshToken(&user)
	if err != nil {
		return domain.RefreshTokenResponse{}, err
	}

	return domain.RefreshTokenResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

// CreateAccessToken implements domain.RefreshTokenUsecase.
func (r *refreshTokenUserCase) CreateAccessToken(user *domain.User) (accessToken string, err error) {
	return r.accesTokenCreator.CreateAccessToken(user)
}

// CreateRefreshToken implements domain.RefreshTokenUsecase.
func (r *refreshTokenUserCase) CreateRefreshToken(user *domain.User) (refreshToken string, err error) {
	return r.refreshTokenCreator.CreateRefreshToken(user)
}

// ExtractIDFromToken implements domain.RefreshTokenUsecase.
func (r *refreshTokenUserCase) ExtractIDFromToken(requestToken string, secret string) (uuid.UUID, error) {
	return r.refreshTokenParser.GetIDFromRefreshToken(requestToken)
}

// GetUserByID implements domain.RefreshTokenUsecase.
func (r *refreshTokenUserCase) GetUserByID(c context.Context, id uuid.UUID) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	user, err := r.userRepository.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
