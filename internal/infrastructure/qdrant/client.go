package qdrant

import (
	"github.com/qdrant/go-client/qdrant"
)

func NewQdrantClient() (*qdrant.Client, error) {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
