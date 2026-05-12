package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	router := gin.Default()
	router.POST("/register", func(ctx *gin.Context) {
		var Acc AuthForm
		if err := ctx.ShouldBindWith(&Acc, binding.JSON); err != nil {
			log.Println("Error: ", err.Error())
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: "Error",
				Data:    nil,
				Success: false,
				Error:   "Internal Server Error",
			})
			return
		}

		newUser := User{
			Id:       idSerial,
			Email:    Acc.Email,
			Password: Acc.Password,
		}
		users = append(users, newUser)
		idSerial++

		log.Printf("Email: %s\nPassword: %s\n", Acc.Email, Acc.Password)
		ctx.JSON(http.StatusOK, Response{
			Message: "Register Success",
			Data:    newUser,
			Success: true,
			Error:   "",
		})
	})

	router.POST("/login", func(ctx *gin.Context) {
		var Acc AuthForm
		if err := ctx.ShouldBindWith(&Acc, binding.JSON); err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: "Error",
				Data:    nil,
				Success: false,
				Error:   "Internal Server Error",
			})
			return
		}

		var isLogin User
		for _, u := range users {
			if u.Email == Acc.Email && u.Password == Acc.Password {
				isLogin = u
				break
			}
		}
		if isLogin.Id == 0 {
			ctx.JSON(http.StatusUnauthorized, Response{
				Message: "Unauthorized",
				Data:    nil,
				Success: false,
				Error:   "Invalid email or password",
			})
			return
		}
		ctx.JSON(http.StatusOK, Response{
			Message: "Login Success",
			Data:    isLogin,
			Success: true,
			Error:   "",
		})
	})
	router.Run("localhost:8081")
}

var idSerial = 1
var users []User

type User struct {
	Id       int
	Email    string
	Password string
}

type Response struct {
	Message string
	Data    any
	Success bool
	Error   string
}

type AuthForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
