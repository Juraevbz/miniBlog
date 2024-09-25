package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthenticateUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = strings.Replace(token, "Bearer ", "", 1)
	claims, err := h.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			Response{Err: errors.Join(errors.New("cannot parse auth token"), err).Error()})
		return
	}

	userID, ok := claims["user_id"]
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			Response{Err: errors.Join(errors.New("cannot parse auth token"), err).Error()})
		return
	}

	c.Set("user_id", userID)
}


type Response struct {
	Err string `json:"error"`
}