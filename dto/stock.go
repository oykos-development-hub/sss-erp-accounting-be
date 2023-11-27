package dto

import (
	"time"

	"gitlab.sudovi.me/erp/accounting-api/data"
)

type StockDTO struct {
	Year               string `json:"year"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Amount             int    `json:"amount"`
	OrganizationUnitID int    `json:"organization_unit_id"`
}

type StockResponseDTO struct {
	ID                 int       `json:"id"`
	Year               string    `json:"year"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	OrganizationUnitID int       `json:"organization_unit_id"`
	Amount             int       `json:"amount"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type StockFilterDTO struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	Year               *string `json:"year"`
	Title              *string `json:"title"`
	Description        *string `json:"description"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
}

func (dto StockDTO) ToStock() *data.Stock {
	return &data.Stock{
		Title:              dto.Title,
		Description:        dto.Description,
		Year:               dto.Year,
		Amount:             dto.Amount,
		OrganizationUnitID: dto.OrganizationUnitID,
	}
}

func ToStockResponseDTO(data data.Stock) StockResponseDTO {
	return StockResponseDTO{
		ID:                 data.ID,
		Title:              data.Title,
		Description:        data.Description,
		Year:               data.Year,
		Amount:             data.Amount,
		OrganizationUnitID: data.OrganizationUnitID,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
	}
}

func ToStockListResponseDTO(stocks []*data.Stock) []StockResponseDTO {
	dtoList := make([]StockResponseDTO, len(stocks))
	for i, x := range stocks {
		dtoList[i] = ToStockResponseDTO(*x)
	}
	return dtoList
}
