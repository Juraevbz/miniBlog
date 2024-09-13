package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	{
		users.POST("/create", h.CreateUser)
		users.GET("/all", h.GetAllUsers)
		users.GET("/:id", h.GetUserByID)
		users.PUT("/update/:id", h.UpdateUser)
		users.DELETE("/delete/:id", h.DeleteUser)
	}

	posts := router.Group("/posts")
	{
		posts.POST("/create", h.CreatePost)
		posts.GET("/:id", h.GetPostByID)
		posts.GET("/all", h.GetAllPosts)
		posts.PUT("/update/:id", h.UpdatePost)
		posts.DELETE("/delete/:id", h.DeletePost)
	}

	comments := router.Group("/comments")
	{
		comments.POST("/create", h.CreateComment)
		comments.GET("/post/:post_id", h.GetComment)
		comments.PUT("/update/:id", h.UpdateComment)
		comments.DELETE("/delete/:id", h.DeleteComment)
	}

	likes := router.Group("/likes")
	{
		likes.POST("/post/:post_id", h.LikePost)
		likes.DELETE("/post/:post_id", h.UnlikePost)
	}

	reposts := router.Group("/reposts")
	{
		reposts.POST("/post/:post_id", h.RepostPost)
		reposts.GET("/user/:id", h.GetUserReposts)
	}

	return router
}
