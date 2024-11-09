package openai

import (
	"context"

	"github.com/sashabaranov/go-openai"
	"github.com/xiao1203/rag_sample/internal/domain/service"
)

type openAIService struct {
	client *openai.Client
}

func NewOpenAIService(apiKey string) service.OpenAIService {
	client := openai.NewClient(apiKey)
	return &openAIService{client: client}
}

func (s *openAIService) VectorizeText(text string) ([]float32, error) {
	resp, err := s.client.CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Model: openai.AdaEmbeddingV2,
			Input: []string{text},
		},
	)
	if err != nil {
		return nil, err
	}
	return resp.Data[0].Embedding, nil
}

func (s *openAIService) GenerateText(prompt string) (string, error) {
	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
