package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/omaraliali1010/go_template/domain"
)

type signupUsecase struct {
	userRepository      domain.UserRepository
	accesTokenCreator   domain.AccessTokenCreator
	refreshTokenCreator domain.RefreshTokenCreator
	timeout             time.Duration
}

func NewSignupUsecase(
	ur domain.UserRepository,
	accesTokenCreator domain.AccessTokenCreator,
	refreshTokenCreator domain.RefreshTokenCreator,
	timeout time.Duration,
) domain.SignupUsecase {
	return &signupUsecase{
		userRepository:      ur,
		accesTokenCreator:   accesTokenCreator,
		refreshTokenCreator: refreshTokenCreator,
		timeout:             timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) (domain.SignupResponse, error) {
	ctx, cancel := context.WithTimeout(c, su.timeout)
	defer cancel()
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return domain.SignupResponse{}, err
	}
	user.Password = hashedPassword
	// Create user
	if err := su.userRepository.CreateUser(ctx, user); err != nil {
		return domain.SignupResponse{}, err
	}

	fetchedUser, err := su.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return domain.SignupResponse{}, err
	}
	// Create access token
	accessToken, err := su.accesTokenCreator.CreateAccessToken(&fetchedUser)
	if err != nil {
		return domain.SignupResponse{}, err
	}

	// Create refresh token
	refreshToken, err := su.refreshTokenCreator.CreateRefreshToken(&fetchedUser)
	if err != nil {
		return domain.SignupResponse{}, err
	}

	// Return response
	return domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func (su *signupUsecase) CreateAccessToken(user *domain.User) (string, error) {
	return su.accesTokenCreator.CreateAccessToken(user)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User) (string, error) {
	return su.refreshTokenCreator.CreateRefreshToken(user)
}
