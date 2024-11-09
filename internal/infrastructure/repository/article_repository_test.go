package repository

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/qdrant/go-client/qdrant"
)

func Test_articleRepository_Save(t *testing.T) {
	// type fields struct {
	// 	client *qdrant.Client
	// }

	client, err := qdrant.NewClient(&qdrant.Config{
		Host:   "localhost",
		Port:   6334,
		UseTLS: false, // 開発環境ではTLSを無効化
	})
	if err != nil {
		t.Fatalf("Qdrantクライアントの作成に失敗しました: %v", err)
	}
	// defer client.Close()
	// Get a context for a minute
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	// Execute health check
	healthCheckResult, err := client.HealthCheck(ctx)
	if err != nil {
		log.Fatalf("Could not get health: %v", err)
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
				texts: &[]string{"テスト文章"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &articleRepository{
				client:         client,
				collectionName: "test",
			}
			if err := a.Save(tt.args.texts); (err != nil) != tt.wantErr {
				t.Errorf("articleRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
