package user

import (
	"context"
	"github.com/jmoiron/sqlx"
	"golearn/internal/models"
	"golearn/internal/repository"
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
	row := transaction.QueryRowContext(ctx, `INSERT INTO users (phone, name, hashed_password) VALUES ($1, $2, $3) returning id;`,
		user.Name, user.SurName, create.Password)

	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Scan, err).Error()
	}
	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.User, cerr.Repository, cerr.Commit, err).Error()
	}
	return id, nil
}

func InitUserRepository(db *sqlx.DB) repository.UserRepo {
	return RepoUser{db: db}
}
