package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/qdrant/go-client/qdrant"
	"github.com/xiao1203/rag_sample/internal/domain/repository"
	"github.com/xiao1203/rag_sample/internal/domain/service"
)

type articleRepository struct {
	client         *qdrant.Client
	openAIService  service.OpenAIService
	collectionName string
}

type SearchResult struct {
	Text  string
	Score float32
}

func NewArticleRepository(client *qdrant.Client, openAIService service.OpenAIService, collectionName string) repository.ArticleRepository {
	return &articleRepository{client: client, openAIService: openAIService, collectionName: collectionName}
}

// FindByID implements repository.ArticleRepository.
func (a *articleRepository) FindSimilarTextsByText(text string, limit uint64) (*[]string, error) {
	// ベクトルの検索

	queryParams := &qdrant.QueryPoints{
		CollectionName: a.collectionName,
		Query:          qdrant.NewQuery(queryVector...),
		Limit:          &limit,
		WithPayload: &qdrant.WithPayloadSelector{
			SelectorOptions: &qdrant.WithPayloadSelector_Enable{
				Enable: true,
			},
		},
	}

	searchResults, err := a.client.Query(context.Background(), queryParams)
	if err != nil {
		return nil, fmt.Errorf("検索に失敗しました: %w", err)
	}

	results := make([]SearchResult, 0, len(searchResults))
	for _, point := range searchResults {
		payload, ok := point.Payload["text"]
		if !ok {
			continue
		}
		text, ok := payload.GetKind().(*qdrant.Value_StringValue)
		if !ok {
			continue
		}
		results = append(results, SearchResult{
			Text:  text.StringValue,
			Score: point.Score,
		})
	}

	texts := make([]string, 0, len(results))
	for _, result := range results {
		texts = append(texts, result.Text)
	}

	return &texts, nil
}

// Save implements repository.ArticleRepository.
func (a *articleRepository) Save(texts *[]string) error {

	// コレクションの存在確認
	// コレクションの存在をチェック
	exists, err := collectionExists(a.client, a.collectionName)
	if err != nil {
		log.Fatalf("コレクションのチェックに失敗しました: %v", err)
	}

	if !exists {
		// コレクションの作成
		err := createCollection(a.client, a.collectionName)
		if err != nil {
			log.Fatalf("コレクションの作成に失敗しました: %v", err)
		}
	}

	// 文章データの登録
	err = insertTextData(a.client, a.collectionName, texts)
	if err != nil {
		log.Fatalf("データの挿入に失敗しました: %v", err)
	}
	return nil
}

func createCollection(client *qdrant.Client, collectionName string) error {
	err := client.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName: collectionName,
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     384, // ベクトルのサイズ（使用するモデルに応じて調整）
			Distance: qdrant.Distance_Cosine,
		}),
	})

	if err != nil {
		return err
	}

	return nil
}

func insertTextData(client *qdrant.Client, collectionName string, texts *[]string) error {
	points := make([]*qdrant.PointStruct, len(*texts))
	for i, text := range *texts {
		// 注意: 実際のアプリケーションでは、ここで文章をベクトル化する処理が必要です
		// この例では、簡単のためにダミーのベクトルを使用しています
		dummyVector := make([]float32, 384)
		for j := range dummyVector {
			dummyVector[j] = float32(i) * 0.1 // ダミーの値
		}

		points[i] = &qdrant.PointStruct{
			Id:      qdrant.NewIDNum(uint64(i + 1)),
			Vectors: qdrant.NewVectors(dummyVector...),
			Payload: qdrant.NewValueMap(map[string]any{"text": text}),
		}
	}

	_, err := client.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points:         points,
	})
	return err
}

func collectionExists(client *qdrant.Client, collectionName string) (bool, error) {
	// コレクションのリストを取得
	collections, err := client.ListCollections(context.Background())
	if err != nil {
		return false, fmt.Errorf("コレクションリストの取得に失敗しました: %w", err)
	}

	// コレクション名が一致するものを探す
	for _, collection := range collections {
		if collection == collectionName {
			return true, nil
		}
	}

	return false, nil
}
