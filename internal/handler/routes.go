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

	user := router.Group("/user")
	{
		user.POST("/sign-up", h.signUp)
		user.POST("/sign-in", h.signIn)
	}

	post := router.Group("/post", h.AuthenticateUser)
	{
		// TODO: implment image saving logic
		post.POST("", h.CreatePost)
		post.GET("/:id", h.GetPostByID)
		post.GET("", h.GetPosts)
		post.PUT("/:id", h.UpdatePost)
		post.DELETE("/delete/:id", h.DeletePost)
	}

	comment := router.Group("/comment", h.AuthenticateUser)
	{
		comment.POST("", h.CreateComment)
		comment.GET("/:id", h.GetCommentByID)
		comment.PUT("/:id", h.UpdateComment)           
		comment.DELETE("/delete/:id", h.DeleteComment)
	}

	like := router.Group("/like", h.AuthenticateUser)
	{
		like.POST("", h.CreateLike)
		like.GET("/:id", h.GetLikeByID)
		like.DELETE("/delete/:id", h.DeleteLike)
	}


	repost := router.Group("/repost", h.AuthenticateUser)
	{
		repost.POST("/:id", h.CreateRepost)             
		repost.GET("/:id", h.GetRepostByID)         
		repost.DELETE("/delete/:id", h.DeleteRepost) 
	}

	return router
}

