package middleware

import (

	"net/http"
	"strings"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

//jwt登录校验
type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {	
	return &LoginJWTMiddlewareBuilder{}
} 

func (l *LoginJWTMiddlewareBuilder)IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)	
	return l
}

// func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if c.Request.Method == "OPTIONS" {
// 			c.Status(http.StatusNoContent)
// 			return
// 		}
// 		for _, path := range l.paths {	
// 			if c.Request.URL.Path == path {
// 				return
// 			}	
// 		}
// 		token:=c.GetHeader("Authorization")
// 		if token==""{
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}
// 		segs:=strings.Split(token,"")
// 		if len(segs)!=2{
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}
// 		tokenStr:=segs[1]
// 		parsedToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 			return []byte("secret"), nil
// 		})
// 		if err!=nil{
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}
// 		if !parsedToken.Valid {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}
// 	}
// }
func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 如果是OPTIONS方法，直接返回204 No Content
        if c.Request.Method == http.MethodOptions {
            c.Status(http.StatusNoContent)
            return
        }

        // 检查是否在忽略路径列表中
        for _, path := range l.paths {
            if c.Request.URL.Path == path {
                return
            }
        }

        // 获取Authorization头
        token := c.GetHeader("Authorization")
        if token == "" {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        // 分割token字符串
        segs := strings.Split(token, " ")
        if len(segs) != 2 {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        // 获取token字符串
        tokenStr := segs[1]

        // 解析JWT
        parsedToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return []byte("secret"), nil
        })
        if err != nil {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        // 检查JWT是否有效
        if !parsedToken.Valid {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }
    }
}

