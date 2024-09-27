package middleware

import (
	"context"
	"strings"

	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/helpers"

	appContext "github.com/BRAVO68WEB/collaborate-with-me/packages/api/utils"

	"github.com/gin-gonic/gin"
)

type JWTServiceImpl struct {
}

func NewJWT() *JWTServiceImpl {
	return &JWTServiceImpl{}
}

func (j *JWTServiceImpl) Auth(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		log := helpers.Logger(ctx)
		if authHeader != "" {
			tokenParts := strings.Split(authHeader, " ")

			if len(tokenParts) == 2 {
				if tokenParts[0] != "Bearer" {
					log.Error("unauthorized")
					return
				}

				tokenString := tokenParts[1]
				if len(tokenString) == 0 {
					log.Error("unauthorized")
					return
				}

				isValid, userClaims := helpers.VerifyJWT(tokenString)

				if !isValid {
					log.Error("unauthorized")
					return
				}

				userID := userClaims.ID

				ctx := appContext.WithUserID(c.Request.Context(), userID)
				c.Request = c.Request.WithContext(ctx)
				c.Next()
			} else {
				log.Error("unauthorized")
				return
			}
		}
	}
}
