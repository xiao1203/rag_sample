package main

import (
	"github.com/labstack/echo/v4"
	"github.com/xiao1203/rag_sample/internal/infrastructure/qdrant"
	"github.com/xiao1203/rag_sample/internal/infrastructure/repository"
	"github.com/xiao1203/rag_sample/internal/interface/controller"
	"github.com/xiao1203/rag_sample/internal/usecase"
	"github.com/xiao1203/rag_sample/router"
)

func main() {
	e := echo.New()

	qdrantClient, err := qdrant.NewQdrantClient()
	if err != nil {
		e.Logger.Fatal(err)
	}

	articleRepo := repository.NewArticleRepository(qdrantClient)
	articleUseCase := usecase.NewArticleUseCase(articleRepo)
	articleController := controller.NewArticleController(articleUseCase)

	router.SetupRoutes(e, articleController)

	e.Logger.Fatal(e.Start(":8080"))
}
