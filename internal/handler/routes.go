package handler

import (
	"mini_blog/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	post := router.Group("/post")
	{
		post.POST("/create", h.CreatePostHandler)
		post.GET("/:id", h.GetPostByIDHandler)
		post.GET("/all", h.GetPostsHandler)
		post.PUT("/update/:id", h.UpdatePostHandler)
		post.DELETE("/delete", h.DeletePostHandler)
	}

	return router
}


