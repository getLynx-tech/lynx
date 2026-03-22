package interfaces

import (
	"context"
	"github.com/getLynx-tech/lynx/internal/domain/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, betterAuthID string) (*entity.User, error)
	GetUserByBetterAuthId(ctx context.Context, betterAuthID string) (*entity.User, error)
}
