package repositories

import (
	"context"
	"database/sql"

	"github.com/osvaldosilitonga/phiraka/server/domain/entity"
)

type userImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) User {
	return &userImpl{
		DB: db,
	}
}

func (u *userImpl) Save(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO users(username, password)
		VALUES ($1, $2)
	`

	_, err := u.DB.ExecContext(ctx, query, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}
