package repository

import "github.com/xiao1203/rag_sample/internal/domain/model"

type ArticleRepository interface {
	Save(article *model.Article) error
	FindByID(id string) (*model.Article, error)
}
