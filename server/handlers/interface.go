package handlers

import (
	"context"

	"github.com/osvaldosilitonga/phiraka/server/domain/web"
)

type User interface {
	Register(ctx context.Context, payload web.RegisterReq) error
}
