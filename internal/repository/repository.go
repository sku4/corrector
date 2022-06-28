package repository

import (
	"context"
	"github.com/sku4/corrector/configs"
	"github.com/sku4/corrector/internal/repository/speller"
	model "github.com/sku4/corrector/models/speller"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository.go

type Corrector interface {
	CheckSpell(texts []string) (model.Response, error)
}

type Repository struct {
	Corrector
}

func NewRepository(ctx context.Context, cfg *configs.Config) *Repository {
	return &Repository{
		Corrector: speller.NewCorrector(ctx, cfg),
	}
}
