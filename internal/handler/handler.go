package handler

import (
	"github.com/SamsonAirapetyan/BWG-test/internal/service"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/invoice", h.Add_into_wallet)
	router.POST("/withdraw", h.Take_from_wallet)
	router.GET("/balance", h.GetAll)

	return router
}
