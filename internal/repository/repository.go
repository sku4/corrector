package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/sku4/corrector/internal/repository/speller"
	model "github.com/sku4/corrector/models/speller"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository.go

type Corrector interface {
	CheckSpell(ctx *gin.Context, texts []string) (model.Response, error)
}

type Repository struct {
	Corrector
}

func NewRepository() *Repository {
	return &Repository{
		Corrector: speller.NewCorrector(),
	}
}
