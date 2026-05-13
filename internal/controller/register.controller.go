package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ilhammursidi/gin-gonic/internal/dto"
)

type RegisterController struct{}

func NewRegisterController() *RegisterController {
	return &RegisterController{}
}

func (r *RegisterController) Register(ctx *gin.Context) {
	var Acc dto.AuthForm
	if err := ctx.ShouldBindWith(&Acc, binding.JSON); err != nil {
		log.Println("Error: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Message: "Error",
			Data:    nil,
			Success: false,
			Error:   "Internal Server Error",
		})
		return
	}

	newUser := dto.User{
		Id:       IdSerial,
		Email:    Acc.Email,
		Password: Acc.Password,
	}
	Users = append(Users, newUser)
	IdSerial++

	if len(newUser.Password) < 7 {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Message: "Unauthorized",
			Data:    nil,
			Success: false,
			Error:   "password minimal 7",
		})
		return
	}

	log.Printf("Email: %s\nPassword: %s\n", Acc.Email, Acc.Password)
	ctx.JSON(http.StatusOK, dto.Response{
		Message: "Register Success",
		Data:    newUser,
		Success: true,
		Error:   "",
	})
}

var IdSerial = 1
var Users []dto.User
