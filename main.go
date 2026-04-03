package main

import (
    "go-hikari/database"
    "go-hikari/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    database.ConnectDB()  

    r := gin.Default()
    routes.SetupRoutes(r) 

    r.Run(":3000")
}