package routes

import (
    "go-hikari/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/blog", controllers.CreateBlog)
    r.GET("/blogs", controllers.GetBlogs)
    r.GET("/blog/:url", controllers.GetBlogByURL)
}