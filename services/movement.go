package services

import (
	"context"

	"gitlab.sudovi.me/erp/accounting-api/data"
	"gitlab.sudovi.me/erp/accounting-api/dto"
	"gitlab.sudovi.me/erp/accounting-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type MovementServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Movement
}

func NewMovementServiceImpl(app *celeritas.Celeritas, repo data.Movement) MovementService {
	return &MovementServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *MovementServiceImpl) CreateMovement(ctx context.Context, input dto.MovementDTO) (*dto.MovementResponseDTO, error) {
	data := input.ToMovement()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToMovementResponseDTO(*data)

	return &res, nil
}

func (h *MovementServiceImpl) UpdateMovement(ctx context.Context, id int, input dto.MovementDTO) (*dto.MovementResponseDTO, error) {
	data := input.ToMovement()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToMovementResponseDTO(*data)

	return &response, nil
}

func (h *MovementServiceImpl) DeleteMovement(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *MovementServiceImpl) GetMovement(id int) (*dto.MovementResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToMovementResponseDTO(*data)

	return &response, nil
}

func (h *MovementServiceImpl) GetMovementList(input *dto.MovementFilterDTO) ([]dto.MovementResponseDTO, *uint64, error) {
	conditionAndExp := &up.AndExpr{}

	if input.RecipientUserID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"recipient_user_id": *input.RecipientUserID})
	}

	if input.OfficeID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"office_id": *input.OfficeID})
	}

	var orders []interface{}

	if input.SortByDateOrder != nil {
		if *input.SortByDateOrder == "asc" {
			orders = append(orders, "-date_order")
		} else {
			orders = append(orders, "date_order")
		}
	}

	orders = append(orders, "-created_at")

	data, total, err := h.repo.GetAll(input.Page, input.Size, conditionAndExp, orders)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToMovementListResponseDTO(data)

	return response, total, nil
}

func (h *MovementServiceImpl) GetMovementReport(input *dto.MovementReportFilterDTO) ([]dto.ArticlesFilterDTO, error) {

	data, err := h.repo.GetAllForReport(input.StartDate, input.EndDate, input.Title, input.OfficeID, input.Exception, input.OrganizationUnitID)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}

	var response []dto.ArticlesFilterDTO

	for _, row := range data {
		response = append(response, dto.ArticlesFilterDTO{
			Year:        row.Year,
			Title:       row.Title,
			Description: row.Description,
			Amount:      row.Amount,
			OfficeID:    row.OfficeID,
		})
	}

	return response, nil
}
