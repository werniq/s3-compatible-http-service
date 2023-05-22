package main

import "github.com/gin-gonic/gin"

func (app *Application) CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Accept", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")

		c.Next()
	}
}
