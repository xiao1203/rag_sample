package repository

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/qdrant/go-client/qdrant"
	"github.com/xiao1203/rag_sample/internal/infrastructure/openai"
)

func Test_articleRepository_Save(t *testing.T) {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host:   "localhost",
		Port:   6334,
		UseTLS: false, // 開発環境ではTLSを無効化
	})
	if err != nil {
		t.Fatalf("Qdrantクライアントの作成に失敗しました: %v", err)
	}

	openAIService := openai.NewOpenAIService(os.Getenv("OPENAI_API_KEY"))

	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// health check
	healthCheckResult, err := client.HealthCheck(ctx)
	if err != nil {
		log.Fatalf("ヘルスチェック失敗: %v", err)
	}
	log.Printf("Qdrant version: %s", healthCheckResult.GetVersion())

	if err != nil {
		t.Fatalf("Qdrantクライアントの作成に失敗しました: %v", err)
	}

	type args struct {
		texts *[]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				texts: &[]string{"テスト文章", "日本の90番目の首相は安倍晋三です。", "日本の92番目の首相は麻生太郎です。"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &articleRepository{
				client:         client,
				openAIService:  openAIService,
				collectionName: "test",
			}
			if err := a.Save(tt.args.texts); (err != nil) != tt.wantErr {
				t.Errorf("articleRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// データ構造体
type TestData struct {
	Text   string    `json:"text"`
	Vector []float32 `json:"vector"`
}

func Test_articleRepository_FindSimilarTextsByText(t *testing.T) {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host:   "localhost",
		Port:   6334,
		UseTLS: false, // 開発環境ではTLSを無効化
	})
	if err != nil {
		t.Fatalf("Qdrantクライアントの作成に失敗しました: %v", err)
	}

	openAIService := openai.NewOpenAIService(os.Getenv("OPENAI_API_KEY"))

	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// health check
	healthCheckResult, err := client.HealthCheck(ctx)
	if err != nil {
		log.Fatalf("ヘルスチェック失敗: %v", err)
	}
	log.Printf("Qdrant version: %s", healthCheckResult.GetVersion())

	if err != nil {
		t.Fatalf("Qdrantクライアントの作成に失敗しました: %v", err)
	}

	a := &articleRepository{
		client:         client,
		openAIService:  openAIService,
		collectionName: "test",
	}

	// // テストデータの登録
	// dataBytes, err := os.ReadFile("./testdata/fixtures/article/find/sample.json")
	// if err != nil {
	// 	t.Fatalf("テストデータの読み込みに失敗しました: %v", err)
	// }

	// var data []TestData
	// if err := json.Unmarshal(dataBytes, &data); err != nil {
	// 	t.Fatalf("JSONデータを読み込むことができません: %v", err)
	// }

	// points := make([]*qdrant.PointStruct, len(data))
	// for i, v := range data {
	// 	vedtor := v.Vector

	// 	points[i] = &qdrant.PointStruct{
	// 		Id:      qdrant.NewIDNum(uint64(i + 1)),
	// 		Vectors: qdrant.NewVectors(vedtor...),
	// 		Payload: qdrant.NewValueMap(map[string]any{"text": v.Text}),
	// 	}
	// }
	// _, err = client.Upsert(context.Background(), &qdrant.UpsertPoints{
	// 	CollectionName: "test",
	// 	Points:         points,
	// })
	// if err != nil {
	// 	t.Fatalf("テストデータの登録に失敗しました: %v", err)
	// }

	type args struct {
		text  string
		limit uint64
	}
	tests := []struct {
		name    string
		args    args
		want    *[]string
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				text:  "日本について教えてください。",
				limit: 5,
			},
			want: &[]string{
				"日本の教育制度は、小学校から高校までが義務教育です。",
				"日本の伝統文化には茶道や華道があり、海外でも注目されています。",
				"日本の伝統的な服装である着物は、特別な行事で着られることが多いです。",
				"日本の電車は時間に正確で、鉄道ネットワークが非常に発達しています。",
				"日本の首都は東京で、人口は約1,400万人です。",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := a.FindSimilarTextsByText(tt.args.text, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("articleRepository.FindSimilarTextsByText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("articleRepository.FindSimilarTextsByText() = %v, want %v", got, tt.want)
			}
		})
	}
}
