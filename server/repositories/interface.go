package repositories

import (
	"context"

	"github.com/osvaldosilitonga/phiraka/server/domain/entity"
)

type User interface {
	Save(ctx context.Context, user *entity.User) error
}
