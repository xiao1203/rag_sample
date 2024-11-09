package service

type OpenAIService interface {
	VectorizeText(text string) ([]float32, error)
	GenerateText(prompt string) (string, error)
}
