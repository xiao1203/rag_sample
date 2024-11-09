package controller

import (
	"net/http"

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
	request, err := newSaveArticleRequest(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if request.URL == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "URL is required"})
	}

	err = c.useCase.SaveArticle(request.URL)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": request.URL})
}

func (c *ArticleController) AnswerQuestion(ctx echo.Context) error {
	request, err := newAnswerQuestionRequest(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if request.Question == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID and question are required"})
	}

	answer, err := c.useCase.AnswerQuestion(request.Question)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(200, map[string]string{"answer": answer})
}
