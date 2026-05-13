package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhammursidi/gin-gonic/internal/controller"
)

func LoginRouter(router *gin.Engine) {
	loginRouter := router.Group("/login")

	controllerLogin := controller.NewLoginController()

	loginRouter.POST("", controllerLogin.Login)
}
