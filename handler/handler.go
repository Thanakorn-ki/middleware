package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct{}

func NewHandler(g *gin.RouterGroup) {
	handler := &handler{}
	g.GET("/health", handler.test)
}

func (h handler) test(c *gin.Context) {
	claims := c.MustGet("claims")
	fmt.Print(claims)
	c.String(http.StatusOK, http.StatusText(http.StatusOK))
}
