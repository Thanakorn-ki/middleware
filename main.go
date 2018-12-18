package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salapao2136/middleware/handler"
	jwt "github.com/salapao2136/middleware/middleware/jwt"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	middle := jwt.NewMiddleware()
	authorized := engine.Group("/")
	authorized.Use(middle.Middleware())
	{
		handler.NewHandler(authorized)
	}

	engine.Run(":8080")
}
