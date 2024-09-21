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
		// TODO: implment image saving logic
		post.POST("", h.CreatePost)
		post.GET("/:id", h.GetPostByID)
		post.GET("", h.GetPosts)
		post.PUT("/:id", h.UpdatePost)
		post.DELETE("/delete/:id", h.DeletePost)
	}

	comment := router.Group("/comment")
	{
		comment.POST("", h.CreateComment)
		comment.GET("/:id", h.GetCommentByID)
		// updateComment
		// deleteComment SOFT
	}

	like := router.Group("/like")
	{
		like.POST("", h.CreateLike)
		like.GET("/:id", h.GetLikeByID)
		// deleteLike HARD
	}

	// TODO: implement Repost
	// createRepost
	// getRepostByID
	// deleteRepost HARD or SOFT deside by yourself 
	return router
}
