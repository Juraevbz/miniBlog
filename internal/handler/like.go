package handler

import (
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateLike(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	in := struct {
		PostID int `json:"post_id"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	like, err := h.service.CreateLike(c, models.Like{
		UserID: int(userID),
		PostID: in.PostID,
	})
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, like)
}

func (h *Handler) GetLikeByID(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	idStr := c.Param("id")
	likeID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	like, err := h.service.GetLikeByID(c, likeID, int(userID))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, like)
}

func (h *Handler) DeleteLike(c *gin.Context) {
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

	err = h.service.DeleteLike(c, id, int(userID))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "like deleted")
}