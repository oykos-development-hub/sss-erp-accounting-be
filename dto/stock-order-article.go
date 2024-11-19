package dto

import (
	"time"

	"gitlab.sudovi.me/erp/accounting-api/data"
)

type StockOrderArticleDTO struct {
	ArticleID int `json:"article_id"`
	StockID   int `json:"stock_id"`
}

type StockOrderArticleResponseDTO struct {
	ID        int       `json:"id"`
	ArticleID int       `json:"article_id"`
	StockID   int       `json:"stock_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (dto StockOrderArticleDTO) ToStockOrderArticle() *data.StockOrderArticle {
	return &data.StockOrderArticle{
		ArticleID: dto.ArticleID,
		StockID:   dto.StockID,
	}
}

func ToStockOrderArticleResponseDTO(data data.StockOrderArticle) StockOrderArticleResponseDTO {
	return StockOrderArticleResponseDTO{
		ID:        data.ID,
		ArticleID: data.ArticleID,
		StockID:   data.StockID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToStockOrderArticleListResponseDTO(stock_order_articles []*data.StockOrderArticle) []StockOrderArticleResponseDTO {
	dtoList := make([]StockOrderArticleResponseDTO, len(stock_order_articles))
	for i, x := range stock_order_articles {
		dtoList[i] = ToStockOrderArticleResponseDTO(*x)
	}
	return dtoList
}
