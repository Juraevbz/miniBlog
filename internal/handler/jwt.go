package handler

import (
	"context"
	"errors"
	"mini_blog/internal/errs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSigningKey = "jfsgvkje9o9309-30"
)

func (h *Handler) GenerateToken(ctx context.Context, claims jwt.MapClaims) (string, error) {
	tokenTTL := time.Duration(12) * time.Hour

	initialClaims := jwt.MapClaims{
		"expires_at": jwt.NewNumericDate(time.Now().UTC().Add(tokenTTL)),
		"issued_at":  jwt.NewNumericDate(time.Now().UTC()),
	}
	for k, v := range claims {
		initialClaims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, initialClaims)

	return token.SignedString([]byte(jwtSigningKey))
}

func (h *Handler) ParseToken(jwtString string) (claims jwt.MapClaims, err error) {
	token, err := jwt.ParseWithClaims(jwtString, &claims, func(token *jwt.Token) (interface{}, error) {
		if sm, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || sm.Name != "HS256" {
			return nil, errors.Join(errs.ErrUnauthorized, err)
		}
		return []byte(jwtSigningKey), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.Join(errs.ErrUnauthorized, err)
	}
	return claims, nil
}
