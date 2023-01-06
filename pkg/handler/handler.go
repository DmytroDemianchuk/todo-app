package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up")
		auth.POST("/sing-in")
	}

	//spiski i ih zadachi
	api := router.Group("/api")
	{
		list := api.Group("/list")
		{
			list.POST("/")
			list.GET("/")
			list.GET("/:id")
			list.PUT("/:id")
			list.DELETE("/:id")

			{
				items := list.Group(":id/items")
				items.POST("/")
				items.GET("/")
				items.GET("/:item_id")
				items.PUT("/:item_id")
				items.DELETE("/:item_id")
			}
		}
	}

	return router
}
