package usecase

import (
	"context"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/omaraliali1010/go_template/domain"
)

type loginUsecase struct {
	userRepository      domain.UserRepository
	accesTokenCreator   domain.AccessTokenCreator
	refreshTokenCreator domain.RefreshTokenCreator
	timeout             time.Duration
}

func NewLoginUsecase(
	ur domain.UserRepository,
	accesTokenCreator domain.AccessTokenCreator,
	refreshTokenCreator domain.RefreshTokenCreator,
	timeout time.Duration,
) domain.LoginUsecase {
	return &loginUsecase{
		userRepository:      ur,
		accesTokenCreator:   accesTokenCreator,
		refreshTokenCreator: refreshTokenCreator,
		timeout:             timeout,
	}
}
func (su *loginUsecase) Login(c context.Context, email string, password string) (domain.LoginResponse, error) {

	log.Println("login_usecase Login: password", password)
	ctx, cancel := context.WithTimeout(c, su.timeout)
	defer cancel()

	user, err := su.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.LoginResponse{}, err
	}
	log.Println("login_usecase Login: user", user)
	isValidPassword := CheckPasswordHash(password, user.Password)
	// Create access token
	accessToken, err := su.accesTokenCreator.CreateAccessToken(&user)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	log.Println("login_usecase Login: accessToken", accessToken)
	// Create refresh token
	refreshToken, err := su.refreshTokenCreator.CreateRefreshToken(&user)
	if err != nil {
		return domain.LoginResponse{}, err
	}
	if !isValidPassword {
		//TODO : return a usefull error that the password in invalid or just a general error
		return domain.LoginResponse{}, nil
	}

	return domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (su *loginUsecase) CreateAccessToken(user *domain.User) (string, error) {
	return su.accesTokenCreator.CreateAccessToken(user)
}

func (su *loginUsecase) CreateRefreshToken(user *domain.User) (string, error) {
	return su.refreshTokenCreator.CreateRefreshToken(user)
}
