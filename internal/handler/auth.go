package handler

import (
	"mini_blog/internal/errs"
	"mini_blog/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) singUp(c *gin.Context) {
	in := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest)
		return
	}

	err := h.service.CreateUser(c.Request.Context(), models.User{
		Username:     in.Username,
		PasswordHash: in.Password,
	})
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, gin.H{"massage": "user created"})
}

func (h *Handler) singIn(c *gin.Context) {
	in := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(400, errs.ErrBadRequest)
		return
	}

	user, err := h.service.Authenticate(c.Request.Context(), in.Username, in.Password)
	if err != nil {
		c.JSON(401, errs.ErrUnauthorized)
		return
	}

	token, err := h.GenerateToken(c, jwt.MapClaims{"user_id": user.ID})
	if err != nil {
		c.JSON(401, errs.ErrUnauthorized)
		return
	}

	c.JSON(200, gin.H{"token": token, "user": user})
}
