package middleware

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}
func (l *LoginMiddlewareBuilder)IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)	
	return l
}
// Build builds a gin middleware that checks if the user is logged in.
// If the request path is /users/login or /users/signup, the middleware
// will not intercept the request. If the user is not logged in, the
// middleware will abort the request with a 401 s tatus code.
func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	gob.Register(time.Time{})
	return func(c *gin.Context) {
	
		if c.Request.URL.Path == "/users/login" || 
			c.Request.URL.Path == "/users/signup" {
			return
		}
		sess := sessions.Default(c)
		id := sess.Get("UserID")
		if id == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return 
		}

		updatetime:=sess.Get("UpdateTime")
		sess.Set("UserID",id)

		sess.Options(sessions.Options{MaxAge: 60 * 60 * 24 * 7})
		now:=time.Now()
		//说明还没有刷新过，刚登陆，还没刷新过
		if updatetime==nil{
			sess.Set("UpdateTime",now)
			sess.Save()
			return
		}
		//updatetime是有的 
		updateTimeValue,_:=updatetime.(time.Time)		
		if now.Sub(updateTimeValue) > time.Second*10 {
			sess.Set("UpdateTime",now)
			sess.Save()
		}
	}
}

func Sub(updateTimeValue time.Time) time.Time {
	panic("unimplemented")
}
