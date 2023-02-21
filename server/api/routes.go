package api

import (
	"github.com/implocell/dashy/core/meme"
	"github.com/labstack/echo/v4"
)

func NewMemeRouter(router *echo.Group, service *meme.MemeService) {
	router.GET("/meme/:id", service.GetMemeByID)
}
