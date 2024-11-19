package handlers

import (
	"net/http"
	"strconv"

	"github.com/oykos-development-hub/celeritas"
	"gitlab.sudovi.me/erp/accounting-api/dto"
	"gitlab.sudovi.me/erp/accounting-api/errors"
	"gitlab.sudovi.me/erp/accounting-api/services"

	"github.com/go-chi/chi/v5"
)

// StockOrderArticleHandler is a concrete type that implements StockOrderArticleHandler
type stockorderarticleHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.StockOrderArticleService
}

// NewStockOrderArticleHandler initializes a new StockOrderArticleHandler with its dependencies
func NewStockOrderArticleHandler(app *celeritas.Celeritas, stockorderarticleService services.StockOrderArticleService) StockOrderArticleHandler {
	return &stockorderarticleHandlerImpl{
		App:     app,
		service: stockorderarticleService,
	}
}

func (h *stockorderarticleHandlerImpl) CreateStockOrderArticle(w http.ResponseWriter, r *http.Request) {
	var input dto.StockOrderArticleDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateStockOrderArticle(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "StockOrderArticle created successfuly", res)
}

func (h *stockorderarticleHandlerImpl) UpdateStockOrderArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.StockOrderArticleDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateStockOrderArticle(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "StockOrderArticle updated successfuly", res)
}

func (h *stockorderarticleHandlerImpl) DeleteStockOrderArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteStockOrderArticle(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "StockOrderArticle deleted successfuly")
}

func (h *stockorderarticleHandlerImpl) GetStockOrderArticleById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetStockOrderArticle(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *stockorderarticleHandlerImpl) GetStockOrderArticleList(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetStockOrderArticleList()
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}
