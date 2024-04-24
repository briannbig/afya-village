package controller

import (
	"log"
	"net/http"

	"github.com/briannbig/afya-village/internal/application/dto"
	"github.com/briannbig/afya-village/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase usecase.UserRegistrationUseCase
}

func NewUserController(usecase usecase.UserRegistrationUseCase) UserController {
	return UserController{usecase: usecase}
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
