package controller

import "github.com/labstack/echo/v4"

type SaveArticleParameter struct {
	URL string `json:"url"`
}

type AnswerQuestionParameter struct {
	Question string `json:"question"`
}

type SaveArticleRequest struct {
	URL string
}

type AnswerQuestionRequest struct {
	Question string
}

func newSaveArticleRequest(ctx echo.Context) (*SaveArticleRequest, error) {
	req := new(SaveArticleRequest)
	if err := ctx.Bind(req); err != nil {
		return nil, err
	}

	return req, nil
}

func newAnswerQuestionRequest(ctx echo.Context) (*AnswerQuestionRequest, error) {
	req := new(AnswerQuestionRequest)
	if err := ctx.Bind(req); err != nil {
		return nil, err
	}

	return req, nil
}
