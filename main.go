package main

import (
	"project-pertama/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	personController := controller.NewPersonController()

	ginEngine := gin.Default()

	ginEngine.GET("/person", personController.GetAll)
	ginEngine.POST("/person", personController.Create)

	err := ginEngine.Run("localhost:8082")
	if err != nil {
		panic(err)
	}
}
