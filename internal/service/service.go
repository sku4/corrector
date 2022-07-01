package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sku4/corrector/internal/repository"
	"github.com/sku4/corrector/internal/service/corrector"
	model "github.com/sku4/corrector/models/corrector"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

type Corrector interface {
	CheckSpell(*gin.Context, model.Request) (model.Response, error)
}

type Service struct {
	Corrector
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Corrector: corrector.NewService(repos.Corrector),
	}
}
