package speller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sku4/corrector/models/speller"
	"github.com/sku4/corrector/pkg/log"
	"io/ioutil"
	"net/http"
	"net/url"
)

const serviceURL = "http://speller.yandex.net/services/spellservice.json/checkTexts"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (r *Client) CheckText(ctx *gin.Context, texts []string) (resp speller.Response, err error) {
	postData := speller.NewRequest()
	postData.Texts = texts
	resp, err = r.send(ctx, *postData)

	return
}

func (r *Client) send(ctx *gin.Context, postData speller.Request) (spellerResp speller.Response, err error) {
	logger := log.LoggerFromGinContext(ctx)

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
