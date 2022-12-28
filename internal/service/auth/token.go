package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"postMaker/internal/entity"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	return s.repo.GetByUsername(ctx, username)
}

func ParseToken(tokenStr string) (*JWTClaim, error) {
	var claims *JWTClaim
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
			return claims, nil
		}
	}

	return claims, err
}
