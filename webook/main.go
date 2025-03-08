package main

import (
	"strings"
	"time"

	"example.com/webook/internal/repository"
	"example.com/webook/internal/repository/dao"
	"example.com/webook/internal/service"
	"example.com/webook/internal/web"
	"example.com/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"

	// "github.com/gin-contrib/sessions/cookie"
	// "github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// main is the entry point of the program.
func main() {
	// Initialize database
	db := initializeDatabase()
	// Initialize web server
	server := initializeWebServer()
	// Initialize user handler
	userHandler := initializeUserHandler(db)
	// Register routes for user handler
	userHandler.RegisterRoutes(server)
	// Run the server
	server.Run(":8080")
}

// initializeDatabase initializes the database.
func initializeDatabase() *gorm.DB {
	// Open database connection
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		// Panic if there is an error
		panic(err)
	}

	// Initialize tables in the database
	if err := dao.InitTables(database); err != nil {
		// Panic if there is an error
		panic(err)
	}
	return database
}

// initializeWebServer initializes the web server.
func initializeWebServer() *gin.Engine {
	// Create a new gin engine
	server := gin.Default()
	// Add CORS middleware
	server.Use(cors.New(cors.Config{
		// Only allow GET, POST and HEAD requests
		AllowMethods:     []string{"GET", "POST", "HEAD"},
		// Only allow Content-Type and Authorization headers
		AllowHeaders:     []string{"Content-Type", "authorization"},
		// Allow credentials
		AllowCredentials: true,
		// Only allow requests from localhost or yourcompany.com
		AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, "http://localhost") || strings.Contains(origin, "yourcompany.com")
		},
		// Set max age to 12 hours
		MaxAge: 12 * time.Hour,
	}))

	// Add session middleware
	// store := cookie.NewStore([]byte("secret"))
	// store := memstore.NewStore([]byte("A#5z!mK4@Nq7$Sv2aPw9^Hx5*Gy3&Jz6"), []byte("X#8z!kL4@Mq7$Rn23Pv9^Hx5*Gy3&Jw6"))
	store,err:=redis.NewStore(16,"tcp","localhost:6379","",
	[]byte("A#5z!mK4@Nq7$Sv2aPw9^Hx5*Gy3&Jz6"), []byte("X#8z!kL4@Mq7$Rn23Pv9^Hx5*Gy3&Jw6"))
	if err!=nil{
		panic(err)
	}
	server.Use(sessions.Sessions("webook", store))
	// Add login middleware
	server.Use(middleware.NewLoginMiddlewareBuilder().
	IgnorePaths("users/login").IgnorePaths("users/signup").Build())
	return server
}

// initializeUserHandler initializes the user handler.
func initializeUserHandler(db *gorm.DB) *web.UserHandler {
	// Create a new user DAO
	userDao := dao.NewUserDao(db)
	// Create a new user repository
	userRepo := repository.NewUserRepository(userDao)
	// Create a new user service
	userService := service.NewUserService(userRepo)
	// Create a new user handler
	return web.NewUserHandler(userService)
}

