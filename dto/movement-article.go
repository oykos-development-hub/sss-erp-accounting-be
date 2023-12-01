package dto

import (
	"time"

	"gitlab.sudovi.me/erp/accounting-api/data"
)

type MovementArticleDTO struct {
	MovementID  int    `json:"movement_id"`
	Year        string `json:"year"`
	Title       string `json:"title"`
	Exception   bool   `json:"exception"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type MovementArticleResponseDTO struct {
	ID          int       `json:"id"`
	MovementID  int       `json:"movement_id"`
	Year        string    `json:"year"`
	Exception   bool      `json:"exception"`
	Title       string    `json:"title"`
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
		Year:        dto.Year,
		Exception:   dto.Exception,
		Title:       dto.Title,
		Description: dto.Description,
		Amount:      dto.Amount,
		MovementID:  dto.MovementID,
	}
}

func ToMovementArticleResponseDTO(data data.MovementArticle) MovementArticleResponseDTO {
	return MovementArticleResponseDTO{
		ID:          data.ID,
		Year:        data.Year,
		Description: data.Description,
		Title:       data.Title,
		Exception:   data.Exception,
		Amount:      data.Amount,
		MovementID:  data.MovementID,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func ToMovementArticleListResponseDTO(movement_articles []*data.MovementArticle) []MovementArticleResponseDTO {
	dtoList := make([]MovementArticleResponseDTO, len(movement_articles))
	for i, x := range movement_articles {
		dtoList[i] = ToMovementArticleResponseDTO(*x)
	}
	return dtoList
}
