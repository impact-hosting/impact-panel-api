package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupControllers(engine *gin.Engine) {
	setupUsersController(engine)
}
