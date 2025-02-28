package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, go")
	})
	server.POST("/post", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, post 方法")
	})
	server.GET("/users/:name", func(ctx *gin.Context) {
		name:=ctx.Param("name")
		ctx.String(http.StatusOK, "hello, 参数路由"+name)
	})
	server.GET("/order", func(ctx *gin.Context) {
		oid:=ctx.Query("id")
		ctx.String(http.StatusOK, "hello, "+oid)
	})
	server.Run(":8080")
}
