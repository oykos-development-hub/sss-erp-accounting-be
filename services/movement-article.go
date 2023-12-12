package services

import (
	"gitlab.sudovi.me/erp/accounting-api/data"
	"gitlab.sudovi.me/erp/accounting-api/dto"
	"gitlab.sudovi.me/erp/accounting-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type MovementArticleServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.MovementArticle
}

func NewMovementArticleServiceImpl(app *celeritas.Celeritas, repo data.MovementArticle) MovementArticleService {
	return &MovementArticleServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *MovementArticleServiceImpl) CreateMovementArticle(input dto.MovementArticleDTO) (*dto.MovementArticleResponseDTO, error) {
	data := input.ToMovementArticle()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToMovementArticleResponseDTO(*data)

	return &res, nil
}

func (h *MovementArticleServiceImpl) UpdateMovementArticle(id int, input dto.MovementArticleDTO) (*dto.MovementArticleResponseDTO, error) {
	data := input.ToMovementArticle()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToMovementArticleResponseDTO(*data)

	return &response, nil
}

func (h *MovementArticleServiceImpl) DeleteMovementArticle(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *MovementArticleServiceImpl) GetMovementArticle(id int) (*dto.MovementArticleResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToMovementArticleResponseDTO(*data)

	return &response, nil
}

func (h *MovementArticleServiceImpl) GetMovementArticleList(input *dto.MovementArticlesFilterDTO) ([]dto.MovementArticleResponseDTO, *uint64, error) {
	conditionAndExp := &up.AndExpr{}

	if input.MovementID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"movement_id": *input.MovementID})
	}

	if input.StockID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"stock_id": *input.StockID})
	}

	data, total, err := h.repo.GetAll(input.Page, input.Size, conditionAndExp)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToMovementArticleListResponseDTO(data)

	return response, total, nil
}
