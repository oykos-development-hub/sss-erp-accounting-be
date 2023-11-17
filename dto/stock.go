package dto

import (
	"time"

	"gitlab.sudovi.me/erp/accounting-api/data"
)

type StockDTO struct {
	ArticleID int `json:"article_id"`
	Amount    int `json:"amount"`
}

type StockResponseDTO struct {
	ID        int       `json:"id"`
	ArticleID int       `json:"article_id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StockFilterDTO struct {
	Page      *int `json:"page"`
	Size      *int `json:"size"`
	ArticleID *int `json:"article_id"`
}

func (dto StockDTO) ToStock() *data.Stock {
	return &data.Stock{
		ArticleID: dto.ArticleID,
		Amount:    dto.Amount,
	}
}

func ToStockResponseDTO(data data.Stock) StockResponseDTO {
	return StockResponseDTO{
		ID:        data.ID,
		ArticleID: data.ArticleID,
		Amount:    data.Amount,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToStockListResponseDTO(stocks []*data.Stock) []StockResponseDTO {
	dtoList := make([]StockResponseDTO, len(stocks))
	for i, x := range stocks {
		dtoList[i] = ToStockResponseDTO(*x)
	}
	return dtoList
}
