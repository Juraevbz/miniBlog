package handler

import (
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePostHandler(c *gin.Context) {
	in := struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, gin.H{"error": errs.ErrBadRequest.Error()})
		return
	}

	post, err := h.service.CreatePostService(c, models.Post{
		Title:    in.Title,
		Content:  in.Content,
		ImageURL: in.ImageURL,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": errs.ErrIntervalServerError.Error()})
		return
	}

	c.JSON(201, post)
}

func (h *Handler) GetPostByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": errs.ErrBadRequest.Error()})
		return
	}

	post, err := h.service.GetPostByIDService(c, postID)
	if err != nil {
		c.JSON(500, gin.H{"error": errs.ErrIntervalServerError.Error()})
		return
	}

	c.JSON(200, gin.H{"post": post})
}

func (h *Handler) GetPostsHandler(c *gin.Context) {
	posts, err := h.service.GetPostsService(c)
	if err != nil {
		c.JSON(500, gin.H{"error": errs.ErrIntervalServerError.Error()})
		return
	}

	c.JSON(200, gin.H{"posts": posts})
}

func (h *Handler) UpdatePostHandler(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": errs.ErrBadRequest.Error()})
		return
	}

	in := struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, gin.H{"error": errs.ErrBadRequest.Error()})
		return
	}

	post, err := h.service.UpdatePostService(c, postID, models.Post{
		Title:    in.Title,
		Content:  in.Content,
		ImageURL: in.ImageURL,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": errs.ErrIntervalServerError.Error()})
		return
	}

	c.JSON(200, gin.H{"success": post})
}

func (h *Handler) DeletePostHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": errs.ErrBadRequest.Error()})
		return
	}

	err = h.service.DeletePostService(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": errs.ErrIntervalServerError})
		return
	}

	c.JSON(200, "post deleted")
}
