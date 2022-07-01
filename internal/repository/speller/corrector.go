package speller

import (
	"context"
	model "github.com/sku4/corrector/models/speller"
	"github.com/sku4/corrector/pkg/speller"
)

type Corrector struct {
	*speller.Client
}

func NewCorrector(ctx context.Context) *Corrector {
	return &Corrector{
		Client: speller.NewClient(ctx),
	}
}

func (c *Corrector) CheckSpell(texts []string) (resp model.Response, err error) {
	return c.Client.CheckText(texts)
}
