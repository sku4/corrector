package corrector

import (
	"context"
	"github.com/sku4/corrector/internal/repository"
	"github.com/sku4/corrector/models/corrector"
	"strings"
)

//go:generate mockgen -source=corrector.go -destination=mocks/corrector.go

type Service struct {
	ctx       context.Context
	corrector repository.Corrector
}

func NewService(ctx context.Context, corrector repository.Corrector) *Service {
	return &Service{
		ctx:       ctx,
		corrector: corrector,
	}
}

func (s *Service) CheckSpell(req corrector.Request) (correctorResp corrector.Response, err error) {
	correctorResp = *corrector.NewResponse()
	spellerResp, err := s.corrector.CheckSpell(req.Texts)

	correctorTexts := make([]string, len(req.Texts))
	for i, text := range req.Texts {
		for _, misspell := range spellerResp[i] {
			if len(misspell.Suggestions) > 0 {
				text = strings.Replace(text, misspell.Word, misspell.Suggestions[0], -1)
			}
		}
		correctorTexts[i] = text
	}
	correctorResp.Texts = correctorTexts

	return
}
