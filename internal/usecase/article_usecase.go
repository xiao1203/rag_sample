package usecase

import (
	"github.com/xiao1203/rag_sample/internal/domain/repository"
	"github.com/xiao1203/rag_sample/internal/domain/service"
	"github.com/xiao1203/rag_sample/internal/infrastructure/webscraper"
)

type ArticleUseCase struct {
	articleRepository repository.ArticleRepository
	scraperService    *webscraper.ScraperService
	openAIService     service.OpenAIService
}

func NewArticleUseCase(
	articleRepository repository.ArticleRepository,
	scraperService *webscraper.ScraperService,
	openAIService service.OpenAIService) *ArticleUseCase {
	return &ArticleUseCase{articleRepository: articleRepository, scraperService: scraperService}
}

func (uc *ArticleUseCase) SaveArticle(url string) error {
	articles, err := uc.scraperService.ExtractText(url)
	if err != nil {
		return err
	}

	err = uc.articleRepository.Save(articles)
	if err != nil {
		return err
	}

	return nil
}

func (uc *ArticleUseCase) AnswerQuestion(id string, question string) (string, error) {
	// 質問に対する参照情報をQdrantから取得
	referenceTexts, err := uc.articleRepository.FindSimilarTextsByText(question, 5)
	if err != nil {
		return "", err
	}

	// 取得した参照情報を改行区切りで整形
	referenceText := ""
	for _, text := range *referenceTexts {
		referenceText += text + "\n"
	}

	// 参照情報を元に質問に対する回答をOpenAIに問い合わせ
	prompt := `
		以下の質問に対し、下記の情報をもとに回答を生成してください。
		[質問]
		` + question + `

		[参照情報]
		` + referenceText + `
	`
	answer, err := uc.openAIService.GenerateText(prompt)
	if err != nil {
		return "", err
	}

	return answer, nil
}
