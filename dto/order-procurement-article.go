package dto

import (
	"time"

	"gitlab.sudovi.me/erp/accounting-api/data"
)

type GetOrderProcurementArticleInputDTO struct {
	Page      *int `json:"page" validate:"omitempty"`
	Size      *int `json:"size" validate:"omitempty"`
	OrderID   *int `json:"order_id"`
	ArticleID *int `json:"article_id"`
}

type OrderProcurementArticleDTO struct {
	OrderID   int `json:"order_id"`
	ArticleID int `json:"article_id"`
	Amount    int `json:"amount"`
}

type OrderProcurementArticleResponseDTO struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"order_id"`
	ArticleID int       `json:"article_id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (dto OrderProcurementArticleDTO) ToOrderProcurementArticle() *data.OrderProcurementArticle {
	return &data.OrderProcurementArticle{
		OrderID:   dto.OrderID,
		ArticleID: dto.ArticleID,
		Amount:    dto.Amount,
	}
}

func ToOrderProcurementArticleResponseDTO(data data.OrderProcurementArticle) OrderProcurementArticleResponseDTO {
	return OrderProcurementArticleResponseDTO{
		ID:        data.ID,
		OrderID:   data.OrderID,
		ArticleID: data.ArticleID,
		Amount:    data.Amount,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToOrderProcurementArticleListResponseDTO(contracts []*data.OrderProcurementArticle) []OrderProcurementArticleResponseDTO {
	dtoList := make([]OrderProcurementArticleResponseDTO, len(contracts))
	for i, x := range contracts {
		dtoList[i] = ToOrderProcurementArticleResponseDTO(*x)
	}
	return dtoList
}
