package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sku4/corrector/models/corrector"
	"net/http"
)

// @Summary Corrector
// @Tags Corrector
// @Description Get answer by webhook corrector command
// @ID corrector-request
// @Accept  json
// @Produce  json
// @Param request body corrector.Request true "Body request"
// @Success 200 {object} corrector.Response
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} corrector.Response
// @Router /corrector [post]
func (h *Handler) correctorRequest(c *gin.Context) {
	var req corrector.Request

	if err := c.BindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.Corrector.CheckSpell(c, req)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
