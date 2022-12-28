package auth

import (
	"context"
	"postMaker/internal/entity"
)

type Repository interface {
	GetByUsername(ctx context.Context, username string) (entity.User, error)
}
