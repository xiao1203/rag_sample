package repository

type ArticleRepository interface {
	Save(*[]string) error
	FindSimilarTextsByText(string, uint64) (*[]string, error)
}
