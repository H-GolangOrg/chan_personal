package main

import (
	"tutorial-go/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	controller "tutorial-go/controller"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "Swagger Example API"

	// localhost:8080/docs/index.html 으로 swagger 확인하기
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1Api := r.Group("/api/v1")
	{
		v1Api.GET("/testGet", controller.ControllerTest)
	}

	r.Run("localhost:8080")
}
