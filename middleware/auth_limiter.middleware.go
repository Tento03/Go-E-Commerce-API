package middleware

import (
	"ecommerce-api/config"
	"ecommerce-api/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginRateLimiter(maxAttempt int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf("rl:login:%s", ip)

		count, err := config.Client.Incr(config.Ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "redis error",
			})
			return
		}

		if count == 1 {
			config.Client.Expire(config.Ctx, key, duration)
		}

		if count > int64(maxAttempt) {
			ttl, _ := config.Client.TTL(config.Ctx, key).Result()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":        "too many login requests attempts",
				"retry_after":  int(ttl.Seconds()),
				"max_attempts": maxAttempt,
			})
			return
		}
		c.Next()
	}
}

func RefreshTokenLimiter(maxAttempt int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		refreshToken, err := c.Cookie("refreshToken")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "refresh token not found",
			})
			return
		}
		hashRT := utils.HashToken(refreshToken)
		key := fmt.Sprintf("rl:refresh:%s:%s", ip, hashRT)

		count, err := config.Client.Incr(config.Ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "redis error",
			})
		}

		if count == 1 {
			config.Client.Expire(config.Ctx, key, duration)
		}

		if count > int64(maxAttempt) {
			ttl, _ := config.Client.TTL(config.Ctx, key).Result()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":        "too many refresh token requests attempts",
				"retry_after":  int(ttl.Seconds()),
				"max_attempts": maxAttempt,
			})
			return
		}
		c.Next()
	}
}
