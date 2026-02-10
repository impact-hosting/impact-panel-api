package controllers

import (
	"impact_api/api/middleware"
	"impact_api/database"
	"impact_api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func setupUsersController(engine *gin.Engine) {
	engine.GET("/users", middleware.AuthMiddleware(), getUsers())
}

func getUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// query params
		limitString := ctx.DefaultQuery("pageSize", "50")
		offsetString := ctx.DefaultQuery("offset", "0")

		limit, err := strconv.Atoi(limitString)
		if err != nil {
			ctx.AbortWithError(500, err)
		}
		offset, err := strconv.Atoi(offsetString)
		if err != nil {
			ctx.AbortWithError(500, err)
		}

		// db query
		results := make([]*models.User, 0)
		result := database.DB.Limit(limit).Offset(offset).Order("created_at DESC").First(&results)
		if result.Error != nil {
			ctx.AbortWithError(500, result.Error)
		}

		ctx.JSON(200, results)
	}
}
