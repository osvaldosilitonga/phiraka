package repositories

import (
	"context"

	"github.com/osvaldosilitonga/phiraka/server/domain/entity"
)

type User interface {
	Save(ctx context.Context, user *entity.User) error
	FindByUsername(ctx context.Context, username string) (entity.User, error)
	FindAllUser(ctx context.Context) ([]entity.User, error)
	UpdateUser(ctx context.Context, username string, user entity.User) error
	DeleteUser(ctx context.Context, username string) error
}
