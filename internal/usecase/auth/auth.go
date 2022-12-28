package auth

import (
	"context"
	"postMaker/internal/service/auth"
)

type UseCase struct {
	auth Auth
}

func NewUseCase(auth Auth) *UseCase {
	return &UseCase{auth: auth}
}

func (cu UseCase) GenerateToken(ctx context.Context, request auth.TokenRequest) (token string, err error) {

	data, err := cu.auth.GetByUsername(ctx, request.Username)
	if err != nil {
		return "", err
	}

	err = auth.CheckPassword(request.Password, data.Password)
	if err != nil {
		return "", err
	}

	tokenString, err := auth.GenerateJWT(data.Id, data.Username)
	if err != nil {
		return "", err
	}

	return tokenString, err
}
