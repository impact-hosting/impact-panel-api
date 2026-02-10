package main

import (
	"fmt"
	"impact_api/api/controllers"
	"impact_api/database"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func main() {
	engine = gin.New()
	engine.Use(gin.Recovery())
	engine.HandleMethodNotAllowed = true
	engine.RedirectFixedPath = true

	controllers.SetupControllers(engine)

	if err := database.ConnectDB(); err != nil {
		panic(err)
	}

	engine.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT")))
}
