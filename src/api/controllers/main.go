package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupControllers(engine *gin.Engine) {
	log.Println("Setting up controllers")
	setupUsersController(engine)
}
