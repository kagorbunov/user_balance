package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kagorbunov/user_balance/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
			users.GET("/", h.updateUser)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUser)
			users.PUT("/trans/:id", h.transUser)
			users.DELETE("/:id", h.deleteUser)

		}
	}
	return router
}
