package controller

import (
	"log"
	"net/http"

	"github.com/briannbig/afya-village/internal/application/dto"
	"github.com/briannbig/afya-village/internal/application/usecase"
	"github.com/briannbig/afya-village/internal/domain/queue"
	"github.com/briannbig/afya-village/internal/infra/database"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase        usecase.UserRegistrationUseCase
}

func NewUserController(db *database.DataBase, q *queue.Queue) UserController {
	userUseCase := usecase.NewUserRegistrationUseCase(database.NewUserRepository(db.Conn), q.RegisterProducer(queue.EventUserCreated))
	return UserController{usecase: *userUseCase}
}

func (u UserController) Register(ctx *gin.Context) {
	var request dto.RequestCreateUser

	if err := ctx.BindJSON(&request); err != nil {
		log.Printf("Could not unmarshal user request: --- %s\n", err.Error())
		return
	}

	user, err := u.usecase.RegisterNewUser(request)
	if err != nil {
		log.Printf("Could not create user: --- %s\n", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, user)

}
