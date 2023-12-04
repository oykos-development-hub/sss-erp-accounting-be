package dto

import (
	"time"

	"gitlab.sudovi.me/erp/accounting-api/data"
)

type MovementDTO struct {
	DateOrder          time.Time `json:"date_order"`
	OrganizationUnitID int       `json:"organization_unit_id"`
	OfficeID           int       `json:"office_id"`
	RecipientUserID    int       `json:"recipient_user_id"`
	Description        string    `json:"description"`
	FileID             int       `json:"file_id"`
}

type MovementResponseDTO struct {
	ID                 int       `json:"id"`
	DateOrder          time.Time `json:"date_order"`
	OrganizationUnitID int       `json:"organization_unit_id"`
	OfficeID           int       `json:"office_id"`
	RecipientUserID    int       `json:"recipient_user_id"`
	Description        string    `json:"description"`
	FileID             int       `json:"file_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type MovementFilterDTO struct {
	Page            *int `json:"page"`
	Size            *int `json:"size"`
	RecipientUserID *int `json:"recipient_user_id"`
	OfficeID        *int `json:"office_id"`
}

type MovementReportFilterDTO struct {
	StartDate          *string `json:"start_date"`
	EndDate            *string `json:"end_date"`
	Title              *string `json:"title"`
	OfficeID           *int    `json:"office_id"`
	Exception          *bool   `json:"exception"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
}

type ArticlesFilterDTO struct {
	Year        string `json:"year"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	OfficeID    int    `json:"office_id"`
}

func (dto MovementDTO) ToMovement() *data.Movement {
	return &data.Movement{
		DateOrder:          dto.DateOrder,
		OrganizationUnitID: dto.OrganizationUnitID,
		OfficeID:           dto.OfficeID,
		RecipientUserID:    dto.RecipientUserID,
		Description:        dto.Description,
		FileID:             dto.FileID,
	}
}

func ToMovementResponseDTO(data data.Movement) MovementResponseDTO {
	return MovementResponseDTO{
		ID:                 data.ID,
		DateOrder:          data.DateOrder,
		OrganizationUnitID: data.OrganizationUnitID,
		OfficeID:           data.OfficeID,
		RecipientUserID:    data.RecipientUserID,
		Description:        data.Description,
		FileID:             data.FileID,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
	}
}

func ToMovementListResponseDTO(movements []*data.Movement) []MovementResponseDTO {
	dtoList := make([]MovementResponseDTO, len(movements))
	for i, x := range movements {
		dtoList[i] = ToMovementResponseDTO(*x)
	}
	return dtoList
}
