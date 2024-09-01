package repository

import (
	"context"
	"learngo/internal/models"
)

type UserRepo interface {
	Create(ctx context.Context, user models.UserCreate) (int, error)
}
