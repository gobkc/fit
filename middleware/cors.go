package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gobkc/fit/conf"
)

func WithCORS() gin.HandlerFunc {
	c := conf.GetConf()
	if c.Cors.Enabled {
		corsConfig := cors.Config{
			AllowOrigins:     c.Cors.AllowedOrigins,
			AllowMethods:     c.Cors.AllowedMethods,
			AllowHeaders:     c.Cors.AllowedHeaders,
			AllowCredentials: c.Cors.AllowCredentials,
			MaxAge:           c.Cors.MaxAge,
			ExposeHeaders:    c.Cors.AllowedHeaders,
		}
		return cors.New(corsConfig)
	}
	return func(c *gin.Context) {
		c.Next()
	}
}
