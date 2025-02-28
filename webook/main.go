package main

import (
	"strings"
	"time"

	"example.com/webook/internal/repository"
	"example.com/webook/internal/repository/dao"
	"example.com/webook/internal/service"
	"example.com/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	db,err:=gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook")) 
	if err != nil { 
		panic(err)
	}

	err=dao.InitTables(db)
	if err != nil {
		panic(err)
	}

	ud:=dao.NewUserDao(db)
	repo:=repository.NewUserRepository(ud)
	svc:=service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	u.RegisterRoutes(server)
	server.Run(":8080")
}