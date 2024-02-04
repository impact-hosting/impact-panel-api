package controllers

import (
	"impact_api/api/middleware"

	"github.com/gin-gonic/gin"
)

func setupUsersController(engine *gin.Engine) {
	engine.GET("/users", middleware.AuthMiddleware(), getUsers())
}

func getUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
