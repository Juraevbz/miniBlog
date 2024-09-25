package handler

import (
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateLike(c *gin.Context) {
	in := struct {
		PostID int `json:"post_id"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	like, err := h.service.CreateLike(c, models.Like{
		PostID: in.PostID,
	})
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, like)
}

func (h *Handler) GetLikeByID(c *gin.Context) {
	idStr := c.Param("id")
	likeID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	like, err := h.service.GetLikeByID(c, likeID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, like)
}

func (h *Handler) DeleteLike(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	err = h.service.DeleteLike(c, id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "like deleted")
}