package services

import (
	"fmt"

	"gitlab.sudovi.me/erp/accounting-api/data"
	"gitlab.sudovi.me/erp/accounting-api/dto"
	"gitlab.sudovi.me/erp/accounting-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type StockServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Stock
}

func NewStockServiceImpl(app *celeritas.Celeritas, repo data.Stock) StockService {
	return &StockServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *StockServiceImpl) CreateStock(input dto.StockDTO) (*dto.StockResponseDTO, error) {
	data := input.ToStock()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToStockResponseDTO(*data)

	return &res, nil
}

func (h *StockServiceImpl) UpdateStock(id int, input dto.StockDTO) (*dto.StockResponseDTO, error) {
	data := input.ToStock()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToStockResponseDTO(*data)

	return &response, nil
}

func (h *StockServiceImpl) DeleteStock(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *StockServiceImpl) GetStock(id int) (*dto.StockResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToStockResponseDTO(*data)

	return &response, nil
}

func (h *StockServiceImpl) GetStockList(input *dto.StockFilterDTO) ([]dto.StockResponseDTO, *uint64, error) {
	conditionAndExp := &up.AndExpr{}

	if input.Title != nil && *input.Title != "" {
		likeCondition := fmt.Sprintf("%%%s%%", *input.Title)

		conditionAndExp = up.And(conditionAndExp, &up.Cond{"title ILIKE": likeCondition})
	}

	data, total, err := h.repo.GetAll(input.Page, input.Size, conditionAndExp)

	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToStockListResponseDTO(data)

	return response, total, nil
}
