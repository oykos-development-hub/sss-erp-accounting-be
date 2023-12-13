package dto

import (
	"time"

	"gitlab.sudovi.me/erp/accounting-api/data"
)

type StockDTO struct {
	Year               string  `json:"year"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	Amount             int     `json:"amount"`
	Exception          bool    `json:"exception"`
	NetPrice           float32 `json:"net_price"`
	VatPercentage      int     `json:"vat_percentage"`
	OrganizationUnitID int     `json:"organization_unit_id"`
}

type StockResponseDTO struct {
	ID                 int       `json:"id"`
	Year               string    `json:"year"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	OrganizationUnitID int       `json:"organization_unit_id"`
	Amount             int       `json:"amount"`
	Exception          bool      `json:"exception"`
	NetPrice           float32   `json:"net_price"`
	VatPercentage      int       `json:"vat_percentage"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type StockFilterDTO struct {
	Page               *int     `json:"page"`
	Size               *int     `json:"size"`
	Year               *string  `json:"year"`
	Title              *string  `json:"title"`
	Description        *string  `json:"description"`
	NetPrice           *float32 `json:"net_price"`
	VatPercentage      *int     `json:"vat_percentage"`
	OrganizationUnitID *int     `json:"organization_unit_id"`
	SortByYear         *string  `json:"sort_by_year"`
	SortByAmount       *string  `json:"sort_by_amount"`
}

func (dto StockDTO) ToStock() *data.Stock {
	return &data.Stock{
		Title:              dto.Title,
		Description:        dto.Description,
		Year:               dto.Year,
		Exception:          dto.Exception,
		Amount:             dto.Amount,
		NetPrice:           dto.NetPrice,
		VatPercentage:      dto.VatPercentage,
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
		Exception:          data.Exception,
		NetPrice:           data.NetPrice,
		VatPercentage:      data.VatPercentage,
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
