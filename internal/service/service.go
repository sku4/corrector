package service

import (
	"context"
	"github.com/sku4/corrector/internal/repository"
	"github.com/sku4/corrector/internal/service/corrector"
	model "github.com/sku4/corrector/models/corrector"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

type Corrector interface {
	CheckSpell(model.Request) (model.Response, error)
}

type Service struct {
	Corrector
}

func NewService(ctx context.Context, repos *repository.Repository) *Service {
	return &Service{
		Corrector: corrector.NewService(ctx, repos.Corrector),
	}
}
