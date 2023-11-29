package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/accounting-api/data"
	"gitlab.sudovi.me/erp/accounting-api/dto"
	"gitlab.sudovi.me/erp/accounting-api/errors"
)

type OrderListServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.OrderList
}

func NewOrderListServiceImpl(app *celeritas.Celeritas, repo data.OrderList) OrderListService {
	return &OrderListServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *OrderListServiceImpl) CreateOrderList(input dto.OrderListDTO) (*dto.OrderListResponseDTO, error) {
	data := input.ToOrderList()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToOrderListResponseDTO(*data)

	return &res, nil
}

func (h *OrderListServiceImpl) UpdateOrderList(id int, input dto.OrderListDTO) (*dto.OrderListResponseDTO, error) {
	data := input.ToOrderList()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToOrderListResponseDTO(*data)

	return &response, nil
}

func (h *OrderListServiceImpl) DeleteOrderList(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *OrderListServiceImpl) GetOrderList(id int) (*dto.OrderListResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToOrderListResponseDTO(*data)

	return &response, nil
}

func (h *OrderListServiceImpl) GetOrderLists(input dto.GetOrderListInputDTO) ([]dto.OrderListResponseDTO, *uint64, error) {
	var (
		combinedCond *up.AndExpr
		conditions   []up.LogicalExpr
	)

	if input.SupplierID != nil {
		cond := up.Cond{
			"supplier_id": input.SupplierID,
		}
		conditions = append(conditions, cond)
	}
	if input.PublicProcurementID != nil {
		cond := up.Cond{
			"public_procurement_id": input.PublicProcurementID,
		}
		conditions = append(conditions, cond)
	}
	if input.OrganizationUnitID != nil {
		cond := up.Cond{
			"organization_unit_id": input.OrganizationUnitID,
		}
		conditions = append(conditions, cond)
	}
	if input.Status != nil {
		cond := up.Cond{
			"status": input.Status,
		}
		conditions = append(conditions, cond)
	}
	if input.Search != nil && *input.Search != "" {
		likeCondition := fmt.Sprintf("%%%s%%", *input.Search)
		searchCond := up.Or(
			up.Cond{"description_recipient ILIKE": likeCondition},
		)
		conditions = append(conditions, searchCond)
	}

	if input.Year != nil && *input.Year != "" {
		year, err := strconv.Atoi(*input.Year)
		if err != nil {
			return nil, nil, err
		}
		startOfYear := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		endOfYear := time.Date(year, time.December, 31, 23, 59, 59, 999999999, time.UTC)

		dateCond := up.And(
			up.Cond{"created_at >=": startOfYear},
			up.Cond{"created_at <=": endOfYear},
		)
		conditions = append(conditions, dateCond)
	}

	if len(conditions) > 0 {
		combinedCond = up.And(conditions...)
	}

	res, total, err := h.repo.GetAll(input.Page, input.Size, combinedCond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToOrderListListResponseDTO(res)

	return response, total, nil
}
