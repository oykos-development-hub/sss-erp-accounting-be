package services

import (
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

	if input.Year != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"year": *input.Year})
	}

	if input.Title != nil {
		titleFilter := *input.Title + "%"
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"title LIKE": titleFilter})
	}

	if input.Description != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"description": *input.Description})
	}

	if input.OrganizationUnitID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"organization_unit_id": *input.OrganizationUnitID})
	}

	var orders []interface{}

	if input.SortByAmount != nil {
		if *input.SortByAmount == "asc" {
			orders = append(orders, "-amount")
		} else {
			orders = append(orders, "amount")
		}
	}

	if input.SortByYear != nil {
		if *input.SortByYear == "asc" {
			orders = append(orders, "-year")
		} else {
			orders = append(orders, "year")
		}
	}

	orders = append(orders, "-created_at")

	data, total, err := h.repo.GetAll(input.Page, input.Size, conditionAndExp, orders)

	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToStockListResponseDTO(data)

	return response, total, nil
}
