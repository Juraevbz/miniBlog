package handler

import (
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	in := struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content" binding:"required"`
		ImageURL string `json:"image_url"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	post, err := h.service.CreatePost(c, models.Post{
		UserID:   int(userID),
		Title:    in.Title,
		Content:  in.Content,
	})
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, post)
}

func (h *Handler) GetPostByID(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	post, err := h.service.GetPostByID(c, postID, int(userID))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"post": post})
}

func (h *Handler) GetPosts(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	postList, err := h.service.GetPosts(c, int(userID))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"post_list": postList})
}

func (h *Handler) UpdatePost(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	in := struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content" binding:"required"`
		ImageURL string `json:"image_url"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	post, err := h.service.UpdatePost(c, postID, models.Post{
		UserID:   int(userID),
		Title:    in.Title,
		Content:  in.Content,
	})
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"updated_post": post})
}

func (h *Handler) DeletePost(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	err = h.service.DeletePost(c, id, int(userID))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "post deleted")
}

