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
	SendToFinance(id int) error
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

type MovementService interface {
	CreateMovement(input dto.MovementDTO) (*dto.MovementResponseDTO, error)
	UpdateMovement(id int, input dto.MovementDTO) (*dto.MovementResponseDTO, error)
	DeleteMovement(id int) error
	GetMovement(id int) (*dto.MovementResponseDTO, error)
	GetMovementList(*dto.MovementFilterDTO) ([]dto.MovementResponseDTO, *uint64, error)
	GetMovementReport(*dto.MovementReportFilterDTO) ([]dto.ArticlesFilterDTO, error)
}

type StockService interface {
	CreateStock(input dto.StockDTO) (*dto.StockResponseDTO, error)
	UpdateStock(id int, input dto.StockDTO) (*dto.StockResponseDTO, error)
	DeleteStock(id int) error
	GetStock(id int) (*dto.StockResponseDTO, error)
	GetStockList(input *dto.StockFilterDTO) ([]dto.StockResponseDTO, *uint64, error)
}

type MovementArticleService interface {
	CreateMovementArticle(input dto.MovementArticleDTO) (*dto.MovementArticleResponseDTO, error)
	UpdateMovementArticle(id int, input dto.MovementArticleDTO) (*dto.MovementArticleResponseDTO, error)
	DeleteMovementArticle(id int) error
	GetMovementArticle(id int) (*dto.MovementArticleResponseDTO, error)
	GetMovementArticleList(input *dto.MovementArticlesFilterDTO) ([]dto.MovementArticleResponseDTO, *uint64, error)
}
