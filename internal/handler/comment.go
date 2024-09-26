package handler

import (
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	in := struct {
		PostID  int    `json:"post_id" binding:"required"`
		Comment string `json:"comment" binding:"required"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	comment, err := h.service.CreateComment(c, models.Comment{
		UserID:  int(userID),
		PostID:  in.PostID,
		Comment: in.Comment,
	})
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(201, comment)
}

func (h *Handler) GetCommentByID(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	idStr := c.Param("id")
	commentID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	comment, err := h.service.GetCommentByID(c, commentID, int(userID))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, comment)
}

func (h *Handler) UpdateComment(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	idStr := c.Param("id")
	commentID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	in := struct {
		Comment string `json:"comment" binding:"required"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	comment, err := h.service.UpdateComment(c, commentID, models.Comment{
		UserID:  int(userID),
		Comment: in.Comment,
	})
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"updated_comment": comment})
}

func (h *Handler) DeleteComment(c *gin.Context) {
	userID := c.GetFloat64("user_id")
	if userID == 0 {
		c.JSON(401, errs.ErrUnauthorized.Error())
		return
	}

	idStr := c.Param("id")
	commentID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	err = h.service.DeleteComment(c, commentID, int(userID))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "comment deleted")
}
