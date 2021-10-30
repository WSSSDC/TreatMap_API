package main

import (
	"treatMap/trick-or-treat/src/data"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/candies", data.Get_Candy)
	router.GET("/users/:id", data.Get_User)
	router.GET("/reports", data.Get_Report)
	router.POST("/users", data.Set_User)
	router.POST("/reports", data.Set_Report)
	router.Run("localhost:8080")
}
