package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ilhammursidi/gin-gonic/internal/dto"
)

type LoginController struct{}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (l *LoginController) Login(ctx *gin.Context) {
	var Acc dto.AuthForm
	if err := ctx.ShouldBindWith(&Acc, binding.JSON); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Message: "Error",
			Data:    nil,
			Success: false,
			Error:   "Internal Server Error",
		})
		return
	}

	var isLogin dto.User
	for _, u := range Users {
		if u.Email == Acc.Email && u.Password == Acc.Password {
			isLogin = u
			break
		}
	}

	if isLogin.Email != Acc.Email || isLogin.Password != Acc.Password {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Message: "Unauthorized",
			Data:    nil,
			Success: false,
			Error:   "Wrong email or password",
		})
		return
	}
	// mending slice containts
	// found := slices.Contains(newUser, isLogin.Email)
	// if !found {
	// 	ctx.JSON(http.StatusUnauthorized, Response{

	// 	})
	// }
	// if isLogin.Id == 0 {
	// 	ctx.JSON(http.StatusUnauthorized, Response{
	// 		Message: "Unauthorized",
	// 		Data:    nil,
	// 		Success: false,
	// 		Error:   "Invalid email or password",
	// 	})
	// 	return
	// }
	ctx.JSON(http.StatusOK, dto.Response{
		Message: "Login Success",
		Data:    isLogin,
		Success: true,
		Error:   "",
	})
}
