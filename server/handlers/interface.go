package handlers

import (
	"context"

	"github.com/osvaldosilitonga/phiraka/server/domain/web"
)

type User interface {
	Register(ctx context.Context, payload web.RegisterReq) error
	Login(ctx context.Context, payload web.LoginReq) (string, error)
	Delete(ctx context.Context, username string) error
	FindAllUser(ctx context.Context) ([]web.FindAllUserResp, error)
}
