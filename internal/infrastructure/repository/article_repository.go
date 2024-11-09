package repository

import (
	"github.com/qdrant/go-client/qdrant"
	"github.com/xiao1203/rag_sample/internal/domain/model"
	"github.com/xiao1203/rag_sample/internal/domain/repository"
)

type articleRepository struct {
	client *qdrant.Client
}

func NewArticleRepository(client *qdrant.Client) repository.ArticleRepository {
	return &articleRepository{client: client}
}

// FindByID implements repository.ArticleRepository.
func (a *articleRepository) FindByID(id string) (*model.Article, error) {
	panic("unimplemented")
}

// Save implements repository.ArticleRepository.
func (a *articleRepository) Save(article *model.Article) error {
	panic("unimplemented")
}
