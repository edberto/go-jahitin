package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"go-jahitin/apipackages"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupAuthMiddleware(toolkit *apipackages.Toolkit) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request

		j := toolkit.AccessAuth

		claims, err := j.ExtractClaims(req)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid Token!"))
			return
		}

		userToken, err := j.GetUserToken(toolkit.DB, claims["access_uuid"].(string))
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Token not found!"))
				return
			default:
				c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Internal Server Error!"))
				return
			}
		}

		if userToken.ExpiredAt.Before(time.Now()) {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Token Expired!"))
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), "userID", userToken.UserID))
		req = req.WithContext(context.WithValue(req.Context(), "accessUUID", userToken.AccessUUID))
		req = req.WithContext(context.WithValue(req.Context(), "role", claims["role"].(string)))

		c.Request = req

		c.Next()
	}
}
