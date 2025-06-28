package jwtservice

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/omaraliali1010/go_template/domain"
)

type JWTService struct {
	AccessSecret       string
	RefreshSecret      string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
}

func (s *JWTService) CreateAccessToken(user *domain.User) (string, error) {
	log.Println("user.ID before token creation:", user.ID)
	claims := &domain.JwtCustomClaims{
		Name: user.Name,
		ID:   user.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.AccessTokenExpiry).Unix(),
		},
	}
	return s.createToken(claims, s.AccessSecret)
}

func (s *JWTService) CreateRefreshToken(user *domain.User) (string, error) {
	claims := &domain.JwtCustomRefreshClaims{
		ID: user.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.RefreshTokenExpiry).Unix(),
		},
	}
	return s.createToken(claims, s.RefreshSecret)
}

func (s *JWTService) createToken(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (s *JWTService) ParseAccessToken(tokenStr string) (*domain.JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.AccessSecret), nil
	})
	log.Println("jwtService parseAccessToken token", token)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*domain.JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	log.Println("ParseAccessToken no error ", claims.ID)
	return claims, nil
}
func (s *JWTService) GetIDFromAccessToken(tokenStr string) (uuid.UUID, error) {
	claims, err := s.ParseAccessToken(tokenStr)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.Parse(claims.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID in token claims: %w", err)
	}

	return id, nil
}

func (s *JWTService) ParseRefreshToken(tokenStr string) (*domain.JwtCustomRefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.JwtCustomRefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.RefreshSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*domain.JwtCustomRefreshClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid refresh token")
	}

	return claims, nil
}
func (s *JWTService) GetIDFromRefreshToken(tokenStr string) (uuid.UUID, error) {
	claims, err := s.ParseRefreshToken(tokenStr)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.Parse(claims.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID in refresh token claims: %w", err)
	}

	return id, nil
}
