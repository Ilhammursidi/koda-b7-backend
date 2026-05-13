package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhammursidi/gin-gonic/internal/controller"
)

func RegisterRouter(router *gin.Engine) {
	registerRouter := router.Group("/register")

	controllerRegister := controller.NewRegisterController()

	registerRouter.POST("", controllerRegister.Register)
}
