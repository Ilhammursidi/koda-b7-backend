package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhammursidi/gin-gonic/internal/router"
)

func main() {
	app := gin.Default()

	router.InitRouter(app)

	app.Run("localhost:8081")
}
