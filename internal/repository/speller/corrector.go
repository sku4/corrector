package speller

import (
	"context"
	"github.com/sku4/corrector/configs"
	model "github.com/sku4/corrector/models/speller"
	"github.com/sku4/corrector/pkg/speller"
)

type Corrector struct {
	*speller.Client
}

func NewCorrector(ctx context.Context, cfg *configs.Config) *Corrector {
	return &Corrector{
		Client: speller.NewClient(ctx, cfg),
	}
}

func (c *Corrector) CheckSpell(texts []string) (resp model.Response, err error) {
	return c.Client.CheckText(texts)
}
