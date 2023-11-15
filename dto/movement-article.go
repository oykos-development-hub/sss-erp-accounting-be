package dto

import (
	"time"

	"gitlab.sudovi.me/erp/accounting-api/data"
)

type MovementArticleDTO struct {
	Title       string `json:"title"`
	MovementID  int    `json:"movement_id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type MovementArticleResponseDTO struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	MovementID  int       `json:"movement_id"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MovementArticlesFilterDTO struct {
	Page       *int `json:"page"`
	Size       *int `json:"size"`
	MovementID *int `json:"movement_id"`
}

func (dto MovementArticleDTO) ToMovementArticle() *data.MovementArticle {
	return &data.MovementArticle{
		Title: dto.Title,
	}
}

func ToMovementArticleResponseDTO(data data.MovementArticle) MovementArticleResponseDTO {
	return MovementArticleResponseDTO{
		ID:        data.ID,
		Title:     data.Title,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToMovementArticleListResponseDTO(movement_articles []*data.MovementArticle) []MovementArticleResponseDTO {
	dtoList := make([]MovementArticleResponseDTO, len(movement_articles))
	for i, x := range movement_articles {
		dtoList[i] = ToMovementArticleResponseDTO(*x)
	}
	return dtoList
}
