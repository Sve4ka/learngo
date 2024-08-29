package repository

import (
	"context"
	"golearn/internal/models"
)

type UserRepo interface {
	Create(ctx context.Context, user models.UserCreate) (int, error)
}
