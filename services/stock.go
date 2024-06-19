package services

import (
	"gitlab.sudovi.me/erp/accounting-api/data"
	"gitlab.sudovi.me/erp/accounting-api/dto"
	newErrors "gitlab.sudovi.me/erp/accounting-api/pkg/errors"

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
		return nil, newErrors.Wrap(err, "repo stock insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo stock get")
	}

	res := dto.ToStockResponseDTO(*data)

	return &res, nil
}

func (h *StockServiceImpl) UpdateStock(id int, input dto.StockDTO) (*dto.StockResponseDTO, error) {
	data := input.ToStock()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo stock get")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo stock get")
	}

	response := dto.ToStockResponseDTO(*data)

	return &response, nil
}

func (h *StockServiceImpl) DeleteStock(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo stock delete")
	}

	return nil
}

func (h *StockServiceImpl) GetStock(id int) (*dto.StockResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo stock get")
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
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"title ILIKE": titleFilter})
	}

	if input.Description != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"description": *input.Description})
	}

	if input.NetPrice != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"net_price": *input.NetPrice})
	}

	if input.VatPercentage != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"vat_percentage": *input.VatPercentage})
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
		return nil, nil, newErrors.Wrap(err, "repo stock get all")
	}
	response := dto.ToStockListResponseDTO(data)

	return response, total, nil
}
