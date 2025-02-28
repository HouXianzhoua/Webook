package main

import (
	"strings"
	"time"

	"example.com/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(func(ctx *gin.Context) {
		println("hello, middleware")
	})
	server.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET","POST"},
		AllowHeaders:     []string{"Content-Type","authorization"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost"){
			// 
			return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge:12*time.Hour,
	}))
	u := web.NewUserHandler()
	u.RegisterRoutes(server)
	server.Run(":8080")
}