package handler

import (
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(c *gin.Context) {
	in := struct {
		PostID  uint   `json:"post_id"`
		Comment string `json:"comment"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	comment, err := h.service.CreateComment(c, models.Comment{
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
	idStr := c.Param("id")
	commentID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	comment, err := h.service.GetCommentByID(c, commentID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, comment)
}

func (h *Handler) UpdateComment(c *gin.Context) {
	idStr := c.Param("id")
	commentID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	in := struct {
		Comment   string `json:"comment"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	comment, err := h.service.UpdateComment(c, commentID, models.Comment{
		Comment: in.Comment,
	})
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"updated_comment": comment})
}

func (h *Handler) DeleteComment(c *gin.Context) {
	idStr := c.Param("id")
	commentID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, errs.ErrBadRequest.Error())
		return
	}

	err = h.service.DeleteComment(c, commentID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "comment deleted")
}

