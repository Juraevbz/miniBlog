package handler

import (
	"mini_blog/internal/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRepost(c *gin.Context) {
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

	repost, err := h.service.CreateRepost(c, postID, int(userID))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, repost)
}

func (h *Handler) GetRepostByID(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	idStr := c.Param("id")
	repostID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	repost, err := h.service.GetRepostByID(c, repostID, int(userID))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, repost)
}

func (h *Handler) DeleteRepost(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	idStr := c.Param("id")
	repostID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	err = h.service.DeleteRepost(c, repostID, int(userID))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "repost deleted")
}

