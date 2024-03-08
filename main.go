package main

import (
	"project-pertama/controller"
	"project-pertama/lib"
	"project-pertama/model"
	"project-pertama/repository"

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

	err = db.AutoMigrate(&model.Person{}, &model.CreditCard{})
	if err != nil {
		panic(err)
	}

	personRepository := repository.NewPersonRepository(db)
	personController := controller.NewPersonController(personRepository)

	ginEngine := gin.Default()

	ginEngine.GET("/person", personController.GetAll)
	ginEngine.POST("/person", personController.Create)
	ginEngine.DELETE("/person/:id", personController.Delete)

	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err = ginEngine.Run("localhost:8082")
	if err != nil {
		panic(err)
	}
}
