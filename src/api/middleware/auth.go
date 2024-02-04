package middleware

import (
	"errors"
	"fmt"
	"impact_api/database"
	"impact_api/models"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// try header auth
		if key := ctx.GetHeader("X-API-Key"); key != "" {
			apiKey := &models.APIKey{}
			res := database.DB.Limit(1).Where("api_key = ?", key).Take(apiKey)
			if res.Error != nil {
				ctx.AbortWithError(500, res.Error)
			}
			if res.RowsAffected != 1 {
				ctx.AbortWithError(401, errors.New("Invalid API Key"))
			}
			// Update API Key `last_used`
			if res = database.DB.Model(&models.APIKey{}).Where("api_key = ?", key).Update("last_used", time.Now().UTC().Nanosecond()); true {
				if res.Error != nil {
					ctx.AbortWithError(500, res.Error)
				}
				if res.RowsAffected != 1 {
					fmt.Printf("Failed to update `last_updated` for API Key '%s'", key)
					ctx.AbortWithError(500, errors.New("An unexpected error has occurred"))
				}
			}
			ctx.Set("user", apiKey.User)
			ctx.Done()
		}
		if sessionId := ctx.GetHeader("SessionID"); sessionId != "" {
			session := &models.Session{}
			res := database.DB.Limit(1).Where("access_token = ?", sessionId).Take(session)
			if res.Error != nil {
				ctx.AbortWithError(500, res.Error)
			}
			if res.RowsAffected != 1 {
				ctx.AbortWithError(401, errors.New("Invalid Session ID"))
			}
			ctx.Set("user", session.User)
			ctx.Done()
		}

		// try cookie auth
		sessionId, err := ctx.Cookie("SessionID")
		if sessionId != "" {
			session := &models.Session{}
			res := database.DB.Limit(1).Where("access_token = ?", sessionId).Take(session)
			if res.Error != nil {
				ctx.AbortWithError(500, res.Error)
			}
			if res.RowsAffected != 1 {
				ctx.AbortWithError(401, errors.New("Invalid Session ID"))
			}
			ctx.Set("user", session.User)
			ctx.Done()
		}
		if err != nil {
			ctx.AbortWithError(401, err)
		}
		ctx.AbortWithError(401, errors.New("No Authorization specified"))
	}
}
