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
	db:=initDB()
	server := initWebserver() 
	u := initUser(db)
	u.RegisterRoutes(server)
	server.Run(":8080")
}

func initWebserver() *gin.Engine {
	server :=gin.Default()   
	server.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type", "authorization"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				//
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDao(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	db,err:=gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook")) 
	if err != nil { 
		panic(err)
	}

	err=dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}