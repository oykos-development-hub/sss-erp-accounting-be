package services

import (
	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/accounting-api/data"
	"gitlab.sudovi.me/erp/accounting-api/dto"
	newErrors "gitlab.sudovi.me/erp/accounting-api/pkg/errors"
)

type OrderProcurementArticleServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.OrderProcurementArticle
}

func NewOrderProcurementArticleServiceImpl(app *celeritas.Celeritas, repo data.OrderProcurementArticle) OrderProcurementArticleService {
	return &OrderProcurementArticleServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *OrderProcurementArticleServiceImpl) CreateOrderProcurementArticle(input dto.OrderProcurementArticleDTO) (*dto.OrderProcurementArticleResponseDTO, error) {
	data := input.ToOrderProcurementArticle()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo order procurement article insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo order procurement article get")
	}

	res := dto.ToOrderProcurementArticleResponseDTO(*data)

	return &res, nil
}

func (h *OrderProcurementArticleServiceImpl) UpdateOrderProcurementArticle(id int, input dto.OrderProcurementArticleDTO) (*dto.OrderProcurementArticleResponseDTO, error) {
	data := input.ToOrderProcurementArticle()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo order procurement article get")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo order procurement article get")
	}

	response := dto.ToOrderProcurementArticleResponseDTO(*data)

	return &response, nil
}

func (h *OrderProcurementArticleServiceImpl) DeleteOrderProcurementArticle(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo order procurement article delete")
	}

	return nil
}

func (h *OrderProcurementArticleServiceImpl) GetOrderProcurementArticle(id int) (*dto.OrderProcurementArticleResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo order procurement article get")
	}
	response := dto.ToOrderProcurementArticleResponseDTO(*data)

	return &response, nil
}

func (h *OrderProcurementArticleServiceImpl) GetOrderProcurementArticles(input dto.GetOrderProcurementArticleInputDTO) ([]dto.OrderProcurementArticleResponseDTO, *uint64, error) {
	var (
		combinedCond *up.AndExpr
		conditions   []up.LogicalExpr
	)

	if input.ArticleID != nil {
		cond := up.Cond{
			"article_id": input.ArticleID,
		}
		conditions = append(conditions, cond)
	}
	if input.OrderID != nil {
		cond := up.Cond{
			"order_id": input.OrderID,
		}
		conditions = append(conditions, cond)
	}

	if len(conditions) > 0 {
		combinedCond = up.And(conditions...)
	}

	res, total, err := h.repo.GetAll(input.Page, input.Size, combinedCond)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "repo order procurement article get all")
	}
	response := dto.ToOrderProcurementArticleListResponseDTO(res)

	return response, total, nil
}
