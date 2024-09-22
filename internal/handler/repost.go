package handler

import (
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRepost(c *gin.Context) {
	in := struct {
		PostID uint `json:"post_id"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	repost, err := h.service.CreateRepost(c, models.Repost{
		PostID: in.PostID,
		UserID: 1,
	})
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, repost)
}

func (h *Handler) GetRepostByID(c *gin.Context) {
	idStr := c.Param("id")
	repostID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	repost, err := h.service.GetRepostByID(c, repostID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, repost)
}

func (h *Handler) DeleteRepost(c *gin.Context) {
	idStr := c.Param("id")
	repostID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	err = h.service.DeleteRepost(c, repostID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "repost deleted")
}