package web

import (
	"fmt"

	"example.com/webook/internal/domain"
	// "example.com/webook/internal/repository"
	"example.com/webook/internal/service"
	"github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserHandler is a handler for user-related requests
type UserHandler struct {
	svc *service.UserService
	emailExp    *regexp2.Regexp
	passwordExp *regexp2.Regexp
}


func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern    = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
		passwordRegexPattern = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$`
	)

	emailExp,err:= regexp2.Compile(emailRegexPattern,0)
	if err != nil {
		panic(err)
	}
	passwordExp, err:= regexp2.Compile(passwordRegexPattern, 0)
	if err != nil {	
		panic(err)	
	}
	return &UserHandler{
		svc:svc,
		emailExp: emailExp,
		passwordExp: passwordExp, 
	}
}

// RegisterRoutes registers all the user-related routes to the given server
func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	// Create a new route group for user operations
	ug := server.Group("/users")
	
	// Define the signup route
	ug.POST("/signup", u.SignUp)
	
	// Define the login route
	ug.POST("/login", u.Login)
	
	// Define the edit route
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignUpReq

	if err := ctx.Bind(&req); err != nil {
		return
	}
	// validate email
	ok, err := u.emailExp.MatchString(req.Email)
	println("validate email")
	if err != nil {
		ctx.String(500, "Internal Server Error")
		return
	}
	if !ok {
		ctx.String(400, "Invalid email")
		return
	}


	// validate password
	println("validate password")
	if req.Password != req.ConfirmPassword {
		ctx.String(400, "Passwords do not match")
		return
	}
	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(500, "Internal Server Error")
		println("Internal Server Error")
		return
	}
	if !ok {
		ctx.String(400, "Invalid password")
		println("Invalid password")
		return
	}

	//
	err=u.svc.Signup(ctx,domain.User{
		Email: req.Email,
		Password: req.Password,
	})
	if err==service.ErrUserDuplicateEmail{
		ctx.String(400, "Duplicate email")
		return
	}
	if err!=nil{
		ctx.String(500, "Internal Server Error")
		return 	
	}


	ctx.String(200, fmt.Sprintf("hello,%v", req))
}

func (u *UserHandler) Login(ctx *gin.Context) {
	// ...
	type LoginReq struct {
		Email    string `json:"email"`	
		Password string `json:"password"`
	}
	 
	var req LoginReq
	if err:=ctx.Bind(&req);err!=nil{
		return
	}
	 
	user,err:=u.svc.Login(ctx,req.Email,req.Password)

	if err==service.ErrInvalidUserOrPassword{
		ctx.String(400, "Invalid user or password")
		return
	} 
	if err != nil {
		ctx.String(500, "Internal Server Error")
		return
	}
	sess:=sessions.Default(ctx)
	sess.Set("userId",user.ID)
	sess.Save() 
	ctx.String(200, fmt.Sprintf("hello,%v", req))
	return
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	// ...
}

func (u *UserHandler) Profile(ctx *gin.Context) {
	// ...
}
