package repositories

import (
	"context"
	"database/sql"
	"errors"

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

func (u *userImpl) FindByUsername(ctx context.Context, username string) (entity.User, error) {
	user := entity.User{}

	query := `
		SELECT id, username, password, create_time
		FROM users
		WHERE username = $1
		LIMIT 1
	`

	err := u.DB.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.Password, &user.CreateTime)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("record not found")
		}

		return user, err
	}

	return user, nil
}

func (u *userImpl) FindAllUser(ctx context.Context) ([]entity.User, error) {
	users := []entity.User{}

	query := `
		SELECT id, username, password, create_time
		FROM users
	`

	rows, err := u.DB.QueryContext(ctx, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return users, errors.New("record not found")
		}

		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		u := entity.User{}

		if err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.CreateTime); err != nil {
			return users, err
		}

		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (u *userImpl) UpdateUser(ctx context.Context, username, password string) error {
	query := `
		UPDATE users
		SET
			username = $1,
			password = $2
		WHERE username = $1
	`

	result, err := u.DB.ExecContext(ctx, query, username, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("record not found")
		}

		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected < 1 {
		return errors.New("no rows affected")
	}

	return nil
}

func (u *userImpl) DeleteUser(ctx context.Context, username string) error {
	query := `
		DELETE FROM users
		WHERE username = $1
	`

	_, err := u.DB.ExecContext(ctx, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("record not found")
		}
		return err
	}

	return nil
}
