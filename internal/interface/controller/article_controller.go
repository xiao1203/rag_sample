package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/xiao1203/rag_sample/internal/usecase"
)

type ArticleController struct {
	useCase *usecase.ArticleUseCase
}

func NewArticleController(useCase *usecase.ArticleUseCase) *ArticleController {
	return &ArticleController{useCase: useCase}
}

func (c *ArticleController) SaveArticle(ctx echo.Context) error {
	url := ctx.FormValue("url")
	if url == "" {
		return ctx.JSON(400, map[string]string{"error": "URL is required"})
	}

	// err := c.useCase.SaveArticle(url)
	// if err != nil {
	// 	return ctx.JSON(500, map[string]string{"error": err.Error()})
	// }

	return ctx.JSON(200, map[string]string{"message": "Article saved successfully"})
}

func (c *ArticleController) AnswerQuestion(ctx echo.Context) error {
	id := ctx.Param("id")
	question := ctx.QueryParam("question")

	if id == "" || question == "" {
		return ctx.JSON(400, map[string]string{"error": "ID and question are required"})
	}

	answer := "answer"
	// answer, err := c.useCase.AnswerQuestion(id, question)
	// if err != nil {
	// 	return ctx.JSON(500, map[string]string{"error": err.Error()})
	// }

	return ctx.JSON(200, map[string]string{"answer": answer})
}
