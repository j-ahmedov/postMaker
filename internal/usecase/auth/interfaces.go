package auth

import (
	"context"
	"postMaker/internal/entity"
)

type Auth interface {
	GetByUsername(ctx context.Context, username string) (entity.User, error)
}
