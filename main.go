package main

import (
	"fmt"
	"os"
	"project-pertama/controller"
	"project-pertama/lib"
	"project-pertama/middleware"
	"project-pertama/model"
	"project-pertama/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "project-pertama/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {

	db, err := lib.InitDatabase()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Person{}, &model.CreditCard{}, &model.User{}, &model.Order{})
	if err != nil {
		panic(err)
	}

	personRepository := repository.NewPersonRepository(db)
	personController := controller.NewPersonController(personRepository)

	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository)

	orderRepository := repository.NewOrderRepository(db)
	orderController := controller.NewOrderController(orderRepository)

	ginEngine := gin.Default()

	ginEngine.POST("/users/register", userController.Register)
	ginEngine.POST("/users/login", userController.Login)

	orderGroup := ginEngine.Group("/orders", middleware.AuthMiddleware)
	orderGroup.GET("", orderController.GetAll)
	orderGroup.POST("", orderController.Create)
	orderGroup.DELETE("/:id", orderController.Delete)

	personGroup := ginEngine.Group("/person", middleware.AuthMiddleware, middleware.AdminMiddleware)
	personGroup.GET("", personController.GetAll)
	personGroup.POST("", personController.Create)
	personGroup.DELETE("/:id", personController.Delete)

	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("PORT")
	port, err := strconv.Atoi(serverPort)
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("%s:%d", serverHost, port)

	err = ginEngine.Run(addr)
	if err != nil {
		panic(err)
	}
}
