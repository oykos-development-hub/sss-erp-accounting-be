package services

import (
	"context"
	"time"

	"gitlab.sudovi.me/erp/accounting-api/dto"
)

type BaseService interface {
	RandomString(n int) string
	Encrypt(text string) (string, error)
	Decrypt(crypto string) (string, error)
}

type OrderListService interface {
	CreateOrderList(ctx context.Context, input dto.OrderListDTO) (*dto.OrderListResponseDTO, error)
	UpdateOrderList(ctx context.Context, id int, input dto.OrderListDTO) (*dto.OrderListResponseDTO, error)
	DeleteOrderList(ctx context.Context, id int) error
	SendToFinance(ctx context.Context, id int) error
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
	CreateMovement(ctx context.Context, input dto.MovementDTO) (*dto.MovementResponseDTO, error)
	UpdateMovement(ctx context.Context, id int, input dto.MovementDTO) (*dto.MovementResponseDTO, error)
	DeleteMovement(ctx context.Context, id int) error
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
	GetAllForReport(date time.Time, organizationUnitID *int) ([]dto.StockResponseDTO, error)
}

type MovementArticleService interface {
	CreateMovementArticle(input dto.MovementArticleDTO) (*dto.MovementArticleResponseDTO, error)
	UpdateMovementArticle(id int, input dto.MovementArticleDTO) (*dto.MovementArticleResponseDTO, error)
	DeleteMovementArticle(id int) error
	GetMovementArticle(id int) (*dto.MovementArticleResponseDTO, error)
	GetMovementArticleList(input *dto.MovementArticlesFilterDTO) ([]dto.MovementArticleResponseDTO, *uint64, error)
}

type LogService interface {
	CreateLog(input dto.LogDTO) (*dto.LogResponseDTO, error)
	UpdateLog(id int, input dto.LogDTO) (*dto.LogResponseDTO, error)
	DeleteLog(id int) error
	GetLog(id int) (*dto.LogResponseDTO, error)
	GetLogList(filter dto.LogFilterDTO) ([]dto.LogResponseDTO, *uint64, error)
}

type ErrorLogService interface {
	CreateErrorLog(err error)
	UpdateErrorLog(id int, input dto.ErrorLogDTO) (*dto.ErrorLogResponseDTO, error)
	DeleteErrorLog(id int) error
	GetErrorLog(id int) (*dto.ErrorLogResponseDTO, error)
	GetErrorLogList(filter dto.ErrorLogFilterDTO) ([]dto.ErrorLogResponseDTO, *uint64, error)
}

type StockOrderArticleService interface {
	CreateStockOrderArticle(input dto.StockOrderArticleDTO) (*dto.StockOrderArticleResponseDTO, error)
	UpdateStockOrderArticle(id int, input dto.StockOrderArticleDTO) (*dto.StockOrderArticleResponseDTO, error)
	DeleteStockOrderArticle(id int) error
	GetStockOrderArticle(id int) (*dto.StockOrderArticleResponseDTO, error)
	GetStockOrderArticleList() ([]dto.StockOrderArticleResponseDTO, error)
}
