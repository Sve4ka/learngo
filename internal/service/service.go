package service

import (
	"context"
	"learngo/internal/models"
)

type UserServ interface {
	Create(ctx context.Context, user models.UserCreate) (int, error)
}
