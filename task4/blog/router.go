package main

import (
	"blog/controller"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// 用户
	r.POST("/api/register", controller.Register) // 注册
	r.POST("/api/login", controller.Login)       // 登录
	// 访客可以使用的功能
	r.GET("/api/posts", controller.ListPost)                 // 文章列表
	r.GET("/api/posts/:id", controller.GetPost)              // 单篇文章
	r.GET("/api/posts/:id/comments", controller.ListComment) // 评论列表
	// 需用户登录的功能
	auth := r.Group("")        // 创建空路径组，方便统一挂中间件
	auth.Use(middleware.JWT()) // 整组先过 JWT 认证
	{
		auth.POST("/api/posts", controller.CreatePost)                 // 发表文章
		auth.PUT("/api/posts/:id", controller.UpdatePost)              // 更新文章
		auth.DELETE("/api/posts/:id", controller.DeletePost)           // 删除文章
		auth.POST("/api/posts/:id/comments", controller.CreateComment) // 发表评论
	}
	return r
}
