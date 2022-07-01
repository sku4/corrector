package speller

import (
	"github.com/gin-gonic/gin"
	model "github.com/sku4/corrector/models/speller"
	"github.com/sku4/corrector/pkg/speller"
)

type Corrector struct {
	*speller.Client
}

func NewCorrector() *Corrector {
	return &Corrector{
		Client: speller.NewClient(),
	}
}

func (c *Corrector) CheckSpell(ctx *gin.Context, texts []string) (resp model.Response, err error) {
	return c.Client.CheckText(ctx, texts)
}
