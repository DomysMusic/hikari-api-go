package controllers

import (
    "context"
    "net/http"
    "go-hikari/database"
    "go-hikari/models"
    "go-hikari/utils"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

func CreateBlog(c *gin.Context) {
    var blog models.Blog
    if err := c.BindJSON(&blog); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    introURL, _ := utils.UploadImage(blog.IntroductionImage)
    challengeURL, _ := utils.UploadImage(blog.ChallengeImage)
    conclusionURL, _ := utils.UploadImage(blog.ConclusionImage)

    blog.IntroductionImage = introURL
    blog.ChallengeImage = challengeURL
    blog.ConclusionImage = conclusionURL

    _, err := database.DB.Collection("blogs").InsertOne(context.Background(), blog)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog created"})
}

func GetBlogs(c *gin.Context) {
    cursor, _ := database.DB.Collection("blogs").Find(context.Background(), bson.M{})
    var blogs []models.Blog
    cursor.All(context.Background(), &blogs)
    c.JSON(http.StatusOK, blogs)
}

func GetBlogByURL(c *gin.Context) {
    url := c.Param("url")
    var blog models.Blog
    err := database.DB.Collection("blogs").FindOne(context.Background(), bson.M{"preferred_url": url}).Decode(&blog)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
        return
    }
    c.JSON(http.StatusOK, blog)
}