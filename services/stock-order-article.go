package services

import (
	"gitlab.sudovi.me/erp/accounting-api/data"
	"gitlab.sudovi.me/erp/accounting-api/dto"
	"gitlab.sudovi.me/erp/accounting-api/errors"

	"github.com/oykos-development-hub/celeritas"
)

type StockOrderArticleServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.StockOrderArticle
}

func NewStockOrderArticleServiceImpl(app *celeritas.Celeritas, repo data.StockOrderArticle) StockOrderArticleService {
	return &StockOrderArticleServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *StockOrderArticleServiceImpl) CreateStockOrderArticle(input dto.StockOrderArticleDTO) (*dto.StockOrderArticleResponseDTO, error) {
	data := input.ToStockOrderArticle()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToStockOrderArticleResponseDTO(*data)

	return &res, nil
}

func (h *StockOrderArticleServiceImpl) UpdateStockOrderArticle(id int, input dto.StockOrderArticleDTO) (*dto.StockOrderArticleResponseDTO, error) {
	data := input.ToStockOrderArticle()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToStockOrderArticleResponseDTO(*data)

	return &response, nil
}

func (h *StockOrderArticleServiceImpl) DeleteStockOrderArticle(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *StockOrderArticleServiceImpl) GetStockOrderArticle(id int) (*dto.StockOrderArticleResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToStockOrderArticleResponseDTO(*data)

	return &response, nil
}

func (h *StockOrderArticleServiceImpl) GetStockOrderArticleList() ([]dto.StockOrderArticleResponseDTO, error) {
	data, err := h.repo.GetAll(nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToStockOrderArticleListResponseDTO(data)

	return response, nil
}
