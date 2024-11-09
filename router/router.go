package router

import (
	"github.com/labstack/echo/v4"
	"github.com/xiao1203/rag_sample/internal/interface/controller"
)

func SetupRoutes(e *echo.Echo, articleController *controller.ArticleController) {
	e.POST("/articles", articleController.SaveArticle)
	e.POST("/rag-answer", articleController.AnswerQuestion)
}
