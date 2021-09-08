package main

import (
	"finansial-service/main/dao"
	"finansial-service/main/endpoints"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	dao.Connect()

	router := gin.Default()
	router.GET("/write-down", endpoints.GetWriteDown)
	router.GET("/addition", endpoints.MakeAddition)
	router.POST("/createUser", endpoints.CreateUser)
	router.GET("/", healthCheck)
	router.GET("/getBalance", endpoints.GetBalance)
	router.POST("/transact", endpoints.MakeTransaction)
	router.NoRoute(onError)
	router.Run("localhost:8080")

}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "App is running successfully",
	})
}

func onError(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Looking page was not found.",
	})
}
