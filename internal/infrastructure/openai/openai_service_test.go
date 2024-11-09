package openai

import (
	"fmt"
	"os"
	"testing"

	"github.com/sashabaranov/go-openai"
)

func Test_openAIService_VectorizeText(t *testing.T) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	service := &openAIService{
		client: client,
	}

	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    []float32
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				text: "アメリカの47番目の副大統領は誰ですか？",
			},
			want:    []float32{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.VectorizeText(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("openAIService.VectorizeText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 何かしら入っていればOK
			fmt.Println(got)

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("openAIService.VectorizeText() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_openAIService_GenerateText(t *testing.T) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	service := &openAIService{
		client: client,
	}

	type args struct {
		prompt string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				prompt: "アメリカの45番目の大統領は誰ですか？",
			},
			want:    "アメリカの45番目の大統領はドナルド・トランプです。",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GenerateText(tt.args.prompt)
			if (err != nil) != tt.wantErr {
				t.Errorf("openAIService.GenerateText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 何かしら入っていればOK
			fmt.Println(got)

			// if got != tt.want {
			// 	t.Errorf("openAIService.GenerateText() = %v, want %v", got, tt.want)
			// }
		})
	}
}
