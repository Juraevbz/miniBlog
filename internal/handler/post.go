package handler

import (
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	in := struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	post, err := h.service.CreatePost(c, models.Post{
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
	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	post, err := h.service.GetPostByID(c, postID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"post": post})
}

func (h *Handler) GetPosts(c *gin.Context) {
	postList, err := h.service.GetPosts(c)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"post_list": postList})
}

func (h *Handler) UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	in := struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	post, err := h.service.UpdatePost(c, postID, models.Post{
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
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	err = h.service.DeletePost(c, id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "post deleted")
}
