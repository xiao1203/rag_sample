package usecase

import (
	"github.com/xiao1203/rag_sample/internal/domain/repository"
)

type ArticleUseCase struct {
	repo repository.ArticleRepository
}

func NewArticleUseCase(repo repository.ArticleRepository) *ArticleUseCase {
	return &ArticleUseCase{repo: repo}
}
