package services

import (
	"gitlab.sudovi.me/erp/accounting-api/dto"
)

type BaseService interface {
	RandomString(n int) string
	Encrypt(text string) (string, error)
	Decrypt(crypto string) (string, error)
}

type OrderListService interface {
	CreateOrderList(input dto.OrderListDTO) (*dto.OrderListResponseDTO, error)
	UpdateOrderList(id int, input dto.OrderListDTO) (*dto.OrderListResponseDTO, error)
	DeleteOrderList(id int) error
	GetOrderList(id int) (*dto.OrderListResponseDTO, error)
	GetOrderLists(data dto.GetOrderListInputDTO) ([]dto.OrderListResponseDTO, *uint64, error)
}

type OrderProcurementArticleService interface {
	CreateOrderProcurementArticle(input dto.OrderProcurementArticleDTO) (*dto.OrderProcurementArticleResponseDTO, error)
	UpdateOrderProcurementArticle(id int, input dto.OrderProcurementArticleDTO) (*dto.OrderProcurementArticleResponseDTO, error)
	DeleteOrderProcurementArticle(id int) error
	GetOrderProcurementArticle(id int) (*dto.OrderProcurementArticleResponseDTO, error)
	GetOrderProcurementArticles(data dto.GetOrderProcurementArticleInputDTO) ([]dto.OrderProcurementArticleResponseDTO, *uint64, error)
}
