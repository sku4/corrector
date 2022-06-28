package speller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sku4/corrector/configs"
	"github.com/sku4/corrector/models/speller"
	"github.com/sku4/corrector/pkg/log"
	"io/ioutil"
	"net/http"
	"net/url"
)

const serviceURL = "http://speller.yandex.net/services/spellservice.json/checkTexts"

type Client struct {
	ctx    context.Context
	config *configs.Config
}

func NewClient(ctx context.Context, cfg *configs.Config) *Client {
	return &Client{
		ctx:    ctx,
		config: cfg,
	}
}

func (r *Client) CheckText(texts []string) (resp speller.Response, err error) {
	postData := speller.NewRequest()
	postData.Texts = texts
	resp, err = r.send(*postData)

	return
}

func (r *Client) send(postData speller.Request) (spellerResp speller.Response, err error) {
	logger := log.LoggerFromContext(r.ctx)

	resp, err := http.PostForm(serviceURL, url.Values{
		"text":   postData.Texts,
		"lang":   {postData.Lang},
		"format": {postData.Format},
	})
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = resp.Body.Close()
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &spellerResp); err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		logger.Info(string(body))
		return spellerResp, errors.New(fmt.Sprint("Response status: ", resp.Status))
	}

	return
}
