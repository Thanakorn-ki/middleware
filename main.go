package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salapao2136/middleware/handler"
	"github.com/salapao2136/middleware/middleware"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	middle := *middleware.NewMiddleware()
	authorized := engine.Group("/")
	authorized.Use(middle.Middleware())
	{
		handler.NewHandler(authorized)
	}

	engine.Run(":8080")
}
