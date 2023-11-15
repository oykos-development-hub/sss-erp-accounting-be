package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
	"gitlab.sudovi.me/erp/accounting-api/dto"
	"gitlab.sudovi.me/erp/accounting-api/errors"
	"gitlab.sudovi.me/erp/accounting-api/services"
)

// OrderProcurementArticleHandler is a concrete type that implements OrderProcurementArticleHandler
type OrderProcurementArticleHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.OrderProcurementArticleService
}

// NewOrderProcurementArticleHandler initializes a new OrderProcurementArticleHandler with its dependencies
func NewOrderProcurementArticleHandler(app *celeritas.Celeritas, OrderProcurementArticleService services.OrderProcurementArticleService) OrderProcurementArticleHandler {
	return &OrderProcurementArticleHandlerImpl{
		App:     app,
		service: OrderProcurementArticleService,
	}
}

func (h *OrderProcurementArticleHandlerImpl) CreateOrderProcurementArticle(w http.ResponseWriter, r *http.Request) {
	var input dto.OrderProcurementArticleDTO
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

	res, err := h.service.CreateOrderProcurementArticle(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "OrderProcurementArticle created successfuly", res)
}

func (h *OrderProcurementArticleHandlerImpl) UpdateOrderProcurementArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.OrderProcurementArticleDTO
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

	res, err := h.service.UpdateOrderProcurementArticle(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "OrderProcurementArticle updated successfuly", res)
}

func (h *OrderProcurementArticleHandlerImpl) DeleteOrderProcurementArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteOrderProcurementArticle(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "OrderProcurementArticle deleted successfuly")
}

func (h *OrderProcurementArticleHandlerImpl) GetOrderProcurementArticleById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetOrderProcurementArticle(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *OrderProcurementArticleHandlerImpl) GetOrderProcurementArticles(w http.ResponseWriter, r *http.Request) {
	var input dto.GetOrderProcurementArticleInputDTO
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

	res, total, err := h.service.GetOrderProcurementArticles(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}
