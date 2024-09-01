package user

import (
	"context"
	"github.com/jmoiron/sqlx"
	"learngo/internal/models"
	"learngo/internal/repository"
)

type RepoUser struct {
	db *sqlx.DB
}

func (repo RepoUser) Create(ctx context.Context, user models.UserCreate) (int, error) {
	var id int
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO users (name, sur_name, email, hashed_password) VALUES ($1, $2, $3, $4) returning id;`,
		user.Name, user.SurName, user.Email, user.PWD)

	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, rbErr
		}
		return 0, err
	}
	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, rbErr
		}
		return 0, err
	}
	return id, nil
}

func InitUserRepository(db *sqlx.DB) repository.UserRepo {
	return RepoUser{db: db}
}
