package handler

import (
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) signUp(c *gin.Context) {
	in := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, errs.ErrBadRequest)
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

	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}

func (h *Handler) signIn(c *gin.Context) {
	type req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var in req
	err := c.BindJSON(&in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Authenticate(c.Request.Context(), in.Username, in.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := h.GenerateToken(c, jwt.MapClaims{"user_id": user.ID})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}