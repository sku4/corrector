package corrector

import (
	"github.com/gin-gonic/gin"
	"github.com/sku4/corrector/internal/repository"
	"github.com/sku4/corrector/models/corrector"
	"strings"
)

//go:generate mockgen -source=corrector.go -destination=mocks/corrector.go

type Service struct {
	corrector repository.Corrector
}

func NewService(corrector repository.Corrector) *Service {
	return &Service{
		corrector: corrector,
	}
}

func (s *Service) CheckSpell(ctx *gin.Context, req corrector.Request) (correctorResp corrector.Response, err error) {
	correctorResp = *corrector.NewResponse()
	spellerResp, err := s.corrector.CheckSpell(ctx, req.Texts)

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
