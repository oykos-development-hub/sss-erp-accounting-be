package dto

import (
	"time"

	"gitlab.sudovi.me/erp/accounting-api/data"
)

type StockDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type StockResponseDTO struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type StockFilterDTO struct {
	Page  *int    `json:"page"`
	Size  *int    `json:"size"`
	Title *string `json:"title"`
}

func (dto StockDTO) ToStock() *data.Stock {
	return &data.Stock{
		Title:       dto.Title,
		Description: dto.Description,
		Amount:      dto.Amount,
	}
}

func ToStockResponseDTO(data data.Stock) StockResponseDTO {
	return StockResponseDTO{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Amount:      data.Amount,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func ToStockListResponseDTO(stocks []*data.Stock) []StockResponseDTO {
	dtoList := make([]StockResponseDTO, len(stocks))
	for i, x := range stocks {
		dtoList[i] = ToStockResponseDTO(*x)
	}
	return dtoList
}
