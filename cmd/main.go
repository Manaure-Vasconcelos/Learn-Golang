package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// inicializa servidor
	server := gin.Default()

	// camada DB connection
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	// camada repository
	productRepository := repository.NewProductRepository(dbConnection)
	// camada usecase
	ProductUseCase := usecase.NewProductUsecase(productRepository)
	//camada controller
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// inicializando a rota chamando controller
	server.GET("/product", ProductController.GetProduct)
	server.POST("/product", ProductController.CreateProduct)
	// inicia servidor
	server.Run(":8080")
}
