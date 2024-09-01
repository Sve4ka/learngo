package user

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"learngo/internal/models"
	"learngo/internal/repository"
	"learngo/internal/service"
	"learngo/pkg/log"
)

type ServUser struct {
	UserRepo repository.UserRepo
	log      *log.Logs
}

func InitUserService(userRepo repository.UserRepo, log *log.Logs) service.UserServ {
	return &ServUser{UserRepo: userRepo, log: log}
}

func (serv ServUser) Create(ctx context.Context, user models.UserCreate) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PWD), 10)
	if err != nil {
		serv.log.Error(err.Error())
		return 0, err
	}
	newUser := models.UserCreate{
		UserBase: user.UserBase,
		PWD:      string(hashedPassword),
	}
	id, err := serv.UserRepo.Create(ctx, newUser)
	if err != nil {
		serv.log.Error(err.Error())
		return 0, err
	}
	serv.log.Info(fmt.Sprintf("create user %v", id))
	return id, nil
}
