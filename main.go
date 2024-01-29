package main

import (
	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func main() {
	engine = gin.New()
	engine.HandleMethodNotAllowed = true
	engine.RedirectFixedPath = true
}
