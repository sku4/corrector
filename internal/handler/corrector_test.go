package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/sku4/corrector/internal/service"
	mock_service "github.com/sku4/corrector/internal/service/mocks"
	"github.com/sku4/corrector/models/corrector"
	testify "github.com/stretchr/testify/assert"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	ListJsonPath = "./corrector_test/list.json"
)

func TestHandler_correctorRequest(t *testing.T) {
	type mockBehavior func(r *mock_service.MockCorrector, req corrector.Request)

	tests := []struct {
		name                 string
		inputBody            string
		inputRequest         corrector.Request
		inputJson            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputJson: ListJsonPath,
			mockBehavior: func(r *mock_service.MockCorrector, req corrector.Request) {
				resp := *corrector.NewResponse()
				r.EXPECT().CheckSpell(req).Return(resp, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"version":"1.0","texts":null}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `[{}]`,
			inputRequest:         corrector.Request{},
			mockBehavior:         func(r *mock_service.MockCorrector, req corrector.Request) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"json: cannot unmarshal array into Go value of type corrector.Request"}`,
		},
		{
			name:      "Service Error",
			inputJson: ListJsonPath,
			mockBehavior: func(r *mock_service.MockCorrector, req corrector.Request) {
				resp := *corrector.NewResponse()
				r.EXPECT().CheckSpell(req).Return(resp, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.inputJson != "" {
				var req corrector.Request
				dataReq, err := os.ReadFile(test.inputJson)
				if err != nil {
					testify.Fail(t, err.Error())
					return
				}
				if err = json.Unmarshal(dataReq, &req); err != nil {
					testify.Fail(t, err.Error())
					return
				}
				test.inputBody = string(dataReq)
				test.inputRequest = req
			}

			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockCorrector(c)
			test.mockBehavior(repo, test.inputRequest)

			services := &service.Service{Corrector: repo}
			handler := Handler{
				ctx:      context.Background(),
				services: *services,
			}

			// Init Endpoint
			r := gin.New()
			r.POST("/corrector", handler.correctorRequest)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/corrector",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
