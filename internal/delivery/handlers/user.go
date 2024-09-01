package handlers

import (
	"context"
	"learngo/internal/models"
	"learngo/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserServ
}

func InitUserHandler(service service.UserServ) UserHandler {
	return UserHandler{
		service: service,
	}
}

// @Summary Create user
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body models.UserCreate true "user create"
// @Success 200 {object} int "Successfully created user"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/create [post]
func (handler UserHandler) CreateUser(g *gin.Context) {
	var newUser models.UserCreate

	if err := g.ShouldBindJSON(&newUser); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := handler.service.Create(ctx, newUser)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}
